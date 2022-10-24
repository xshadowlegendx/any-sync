package syncservice

import (
	"context"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/spacesyncproto"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/net/rpc/rpctest"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/consensus/consensusproto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/require"
	"storj.io/drpc"
	"testing"
	"time"
)

type testPeer struct {
	id string
	drpc.Conn
}

func (t testPeer) Id() string {
	return t.id
}

func (t testPeer) LastUsage() time.Time {
	return time.Now()
}

func (t testPeer) UpdateLastUsage() {}

type testServer struct {
	stream        chan spacesyncproto.DRPCSpace_StreamStream
	addLog        func(ctx context.Context, req *consensusproto.AddLogRequest) error
	addRecord     func(ctx context.Context, req *consensusproto.AddRecordRequest) error
	releaseStream chan error
	watchErrOnce  bool
}

func (t *testServer) HeadSync(ctx context.Context, request *spacesyncproto.HeadSyncRequest) (*spacesyncproto.HeadSyncResponse, error) {
	panic("implement me")
}

func (t *testServer) PushSpace(ctx context.Context, request *spacesyncproto.PushSpaceRequest) (*spacesyncproto.PushSpaceResponse, error) {
	panic("implement me")
}

func (t *testServer) Stream(stream spacesyncproto.DRPCSpace_StreamStream) error {
	t.stream <- stream
	return <-t.releaseStream
}

func (t *testServer) waitStream(test *testing.T) spacesyncproto.DRPCSpace_StreamStream {
	select {
	case <-time.After(time.Second * 5):
		test.Fatalf("waiteStream timeout")
	case st := <-t.stream:
		return st
	}
	return nil
}

type fixture struct {
	testServer   *testServer
	drpcTS       *rpctest.TesServer
	client       spacesyncproto.DRPCSpaceClient
	clientStream spacesyncproto.DRPCSpace_StreamStream
	serverStream spacesyncproto.DRPCSpace_StreamStream
	pool         *streamPool
	localId      peer.ID
	remoteId     peer.ID
}

func newFixture(t *testing.T, localId, remoteId peer.ID, handler MessageHandler) *fixture {
	fx := &fixture{
		testServer: &testServer{},
		drpcTS:     rpctest.NewTestServer(),
		localId:    localId,
		remoteId:   remoteId,
	}
	fx.testServer.stream = make(chan spacesyncproto.DRPCSpace_StreamStream, 1)
	require.NoError(t, spacesyncproto.DRPCRegisterSpace(fx.drpcTS.Mux, fx.testServer))
	clientWrapper := rpctest.NewSecConnWrapper(nil, nil, localId, remoteId)
	p := &testPeer{id: localId.String(), Conn: fx.drpcTS.DialWrapConn(nil, clientWrapper)}
	fx.client = spacesyncproto.NewDRPCSpaceClient(p)

	var err error
	fx.clientStream, err = fx.client.Stream(context.Background())
	require.NoError(t, err)
	fx.serverStream = fx.testServer.waitStream(t)
	fx.pool = newStreamPool(handler).(*streamPool)

	return fx
}

func (fx *fixture) run(t *testing.T) chan error {
	waitCh := make(chan error)
	go func() {
		err := fx.pool.AddAndReadStreamSync(fx.clientStream)
		waitCh <- err
	}()

	time.Sleep(time.Millisecond * 10)
	fx.pool.Lock()
	require.Equal(t, fx.pool.peerStreams[fx.remoteId.String()], fx.clientStream)
	fx.pool.Unlock()

	return waitCh
}

func TestStreamPool_AddAndReadStreamAsync(t *testing.T) {
	remId := peer.ID("remoteId")

	t.Run("client close", func(t *testing.T) {
		fx := newFixture(t, "", remId, nil)
		waitCh := fx.run(t)

		err := fx.clientStream.Close()
		require.NoError(t, err)
		err = <-waitCh

		require.Error(t, err)
		require.Nil(t, fx.pool.peerStreams[remId.String()])
	})
	t.Run("server close", func(t *testing.T) {
		fx := newFixture(t, "", remId, nil)
		waitCh := fx.run(t)

		err := fx.serverStream.Close()
		require.NoError(t, err)

		err = <-waitCh
		require.Error(t, err)
		require.Nil(t, fx.pool.peerStreams[remId.String()])
	})
}

func TestStreamPool_Close(t *testing.T) {
	remId := peer.ID("remoteId")

	t.Run("client close", func(t *testing.T) {
		fx := newFixture(t, "", remId, nil)
		fx.run(t)
		var events []string
		recvChan := make(chan struct{})
		go func() {
			fx.pool.Close()
			events = append(events, "pool_close")
			recvChan <- struct{}{}
		}()
		time.Sleep(50 * time.Millisecond) //err = <-waitCh
		events = append(events, "stream_close")
		err := fx.clientStream.Close()
		require.NoError(t, err)
		<-recvChan
		require.Equal(t, []string{"stream_close", "pool_close"}, events)
	})
	t.Run("server close", func(t *testing.T) {
		fx := newFixture(t, "", remId, nil)
		fx.run(t)
		var events []string
		recvChan := make(chan struct{})
		go func() {
			fx.pool.Close()
			events = append(events, "pool_close")
			recvChan <- struct{}{}
		}()
		time.Sleep(50 * time.Millisecond) //err = <-waitCh
		events = append(events, "stream_close")
		err := fx.clientStream.Close()
		require.NoError(t, err)
		<-recvChan
		require.Equal(t, []string{"stream_close", "pool_close"}, events)
	})
}

func TestStreamPool_ReceiveMessage(t *testing.T) {
	remId := peer.ID("remoteId")
	t.Run("pool receive message from server", func(t *testing.T) {
		objectId := "objectId"
		msg := &spacesyncproto.ObjectSyncMessage{
			ObjectId: objectId,
		}
		recvChan := make(chan struct{})
		fx := newFixture(t, "", remId, func(ctx context.Context, senderId string, message *spacesyncproto.ObjectSyncMessage) (err error) {
			require.Equal(t, msg, message)
			recvChan <- struct{}{}
			return nil
		})
		waitCh := fx.run(t)

		err := fx.serverStream.Send(msg)
		require.NoError(t, err)
		<-recvChan
		err = fx.clientStream.Close()
		require.NoError(t, err)
		err = <-waitCh

		require.Error(t, err)
		require.Nil(t, fx.pool.peerStreams[remId.String()])
	})
}

func TestStreamPool_HasActiveStream(t *testing.T) {
	remId := peer.ID("remoteId")
	t.Run("pool has active stream", func(t *testing.T) {
		fx := newFixture(t, "", remId, nil)
		waitCh := fx.run(t)
		require.True(t, fx.pool.HasActiveStream(remId.String()))

		err := fx.clientStream.Close()
		require.NoError(t, err)
		err = <-waitCh

		require.Error(t, err)
		require.Nil(t, fx.pool.peerStreams[remId.String()])
	})
	t.Run("pool has no active stream", func(t *testing.T) {
		fx := newFixture(t, "", remId, nil)
		waitCh := fx.run(t)
		err := fx.clientStream.Close()
		require.NoError(t, err)
		err = <-waitCh
		require.Error(t, err)
		require.False(t, fx.pool.HasActiveStream(remId.String()))
		require.Nil(t, fx.pool.peerStreams[remId.String()])
	})
}

func TestStreamPool_SendAsync(t *testing.T) {
	remId := peer.ID("remoteId")
	t.Run("pool send async to server", func(t *testing.T) {
		objectId := "objectId"
		msg := &spacesyncproto.ObjectSyncMessage{
			ObjectId: objectId,
		}
		fx := newFixture(t, "", remId, nil)
		recvChan := make(chan struct{})
		go func() {
			message, err := fx.serverStream.Recv()
			require.NoError(t, err)
			require.Equal(t, msg, message)
			recvChan <- struct{}{}
		}()
		waitCh := fx.run(t)

		err := fx.pool.SendAsync([]string{remId.String()}, msg)
		require.NoError(t, err)
		<-recvChan
		err = fx.clientStream.Close()
		require.NoError(t, err)
		err = <-waitCh

		require.Error(t, err)
		require.Nil(t, fx.pool.peerStreams[remId.String()])
	})
}

func TestStreamPool_SendSync(t *testing.T) {
	remId := peer.ID("remoteId")
	t.Run("pool send sync to server", func(t *testing.T) {
		objectId := "objectId"
		payload := []byte("payload")
		msg := &spacesyncproto.ObjectSyncMessage{
			ObjectId: objectId,
		}
		fx := newFixture(t, "", remId, nil)
		go func() {
			message, err := fx.serverStream.Recv()
			require.NoError(t, err)
			require.Equal(t, msg.ObjectId, message.ObjectId)
			require.NotEmpty(t, message.ReplyId)
			message.Payload = payload
			err = fx.serverStream.Send(message)
			require.NoError(t, err)
		}()
		waitCh := fx.run(t)
		res, err := fx.pool.SendSync(remId.String(), msg)
		require.NoError(t, err)
		require.Equal(t, payload, res.Payload)
		err = fx.clientStream.Close()
		require.NoError(t, err)
		err = <-waitCh

		require.Error(t, err)
		require.Nil(t, fx.pool.peerStreams[remId.String()])
	})

	t.Run("pool send sync timeout", func(t *testing.T) {
		objectId := "objectId"
		msg := &spacesyncproto.ObjectSyncMessage{
			ObjectId: objectId,
		}
		fx := newFixture(t, "", remId, nil)
		syncWaitPeriod = time.Millisecond * 30
		go func() {
			message, err := fx.serverStream.Recv()
			require.NoError(t, err)
			require.Equal(t, msg.ObjectId, message.ObjectId)
			require.NotEmpty(t, message.ReplyId)
		}()
		waitCh := fx.run(t)
		_, err := fx.pool.SendSync(remId.String(), msg)
		require.Equal(t, ErrSyncTimeout, err)
		err = fx.clientStream.Close()
		require.NoError(t, err)
		err = <-waitCh

		require.Error(t, err)
		require.Nil(t, fx.pool.peerStreams[remId.String()])
	})
}

func TestStreamPool_BroadcastAsync(t *testing.T) {
	remId := peer.ID("remoteId")
	t.Run("pool broadcast async to server", func(t *testing.T) {
		objectId := "objectId"
		msg := &spacesyncproto.ObjectSyncMessage{
			ObjectId: objectId,
		}
		fx := newFixture(t, "", remId, nil)
		recvChan := make(chan struct{})
		go func() {
			message, err := fx.serverStream.Recv()
			require.NoError(t, err)
			require.Equal(t, msg, message)
			recvChan <- struct{}{}
		}()
		waitCh := fx.run(t)

		err := fx.pool.BroadcastAsync(msg)
		require.NoError(t, err)
		<-recvChan
		err = fx.clientStream.Close()
		require.NoError(t, err)
		err = <-waitCh

		require.Error(t, err)
		require.Nil(t, fx.pool.peerStreams[remId.String()])
	})
}