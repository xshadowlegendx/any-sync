package syncservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/spacesyncproto"
	"github.com/libp2p/go-libp2p-core/sec"
	"storj.io/drpc/drpcctx"
	"sync"
	"sync/atomic"
)

var ErrEmptyPeer = errors.New("don't have such a peer")
var ErrStreamClosed = errors.New("stream is already closed")

const maxSimultaneousOperationsPerStream = 10

// StreamPool can be made generic to work with different streams
type StreamPool interface {
	SyncClient
	AddAndReadStreamSync(stream spacesyncproto.SpaceStream) (err error)
	AddAndReadStreamAsync(stream spacesyncproto.SpaceStream)
	HasStream(peerId string) bool
	Close() (err error)
}

type SyncClient interface {
	SendSync(peerId string, message *spacesyncproto.ObjectSyncMessage) (reply *spacesyncproto.ObjectSyncMessage, err error)
	SendAsync(peerId string, message *spacesyncproto.ObjectSyncMessage) (err error)
	BroadcastAsync(message *spacesyncproto.ObjectSyncMessage) (err error)
}

type MessageHandler func(ctx context.Context, senderId string, message *spacesyncproto.ObjectSyncMessage) (err error)

type responseWaiter struct {
	ch chan *spacesyncproto.ObjectSyncMessage
}

type streamPool struct {
	sync.Mutex
	peerStreams    map[string]spacesyncproto.SpaceStream
	messageHandler MessageHandler
	wg             *sync.WaitGroup
	waiters        map[string]responseWaiter
	waitersMx      sync.Mutex
	counter        uint64
}

func newStreamPool(messageHandler MessageHandler) StreamPool {
	return &streamPool{
		peerStreams:    make(map[string]spacesyncproto.SpaceStream),
		messageHandler: messageHandler,
		wg:             &sync.WaitGroup{},
	}
}

func (s *streamPool) HasStream(peerId string) (res bool) {
	_, err := s.getStream(peerId)
	return err == nil
}

func (s *streamPool) SendSync(
	peerId string,
	msg *spacesyncproto.ObjectSyncMessage) (reply *spacesyncproto.ObjectSyncMessage, err error) {
	newCounter := atomic.AddUint64(&s.counter, 1)
	msg.TrackingId = genStreamPoolKey(peerId, msg.TreeId, newCounter)

	s.waitersMx.Lock()
	waiter := responseWaiter{
		ch: make(chan *spacesyncproto.ObjectSyncMessage),
	}
	s.waiters[msg.TrackingId] = waiter
	s.waitersMx.Unlock()

	err = s.SendAsync(peerId, msg)
	if err != nil {
		return
	}

	reply = <-waiter.ch
	return
}

func (s *streamPool) SendAsync(peerId string, message *spacesyncproto.ObjectSyncMessage) (err error) {
	stream, err := s.getStream(peerId)
	if err != nil {
		return
	}

	return stream.Send(message)
}

func (s *streamPool) getStream(id string) (stream spacesyncproto.SpaceStream, err error) {
	s.Lock()
	defer s.Unlock()
	stream, exists := s.peerStreams[id]
	if !exists {
		err = ErrEmptyPeer
		return
	}

	select {
	case <-stream.Context().Done():
		delete(s.peerStreams, id)
		err = ErrStreamClosed
	default:
	}

	return
}

func (s *streamPool) getAllStreams() (streams []spacesyncproto.SpaceStream) {
	s.Lock()
	defer s.Unlock()
Loop:
	for id, stream := range s.peerStreams {
		select {
		case <-stream.Context().Done():
			delete(s.peerStreams, id)
			continue Loop
		default:
		}
		streams = append(streams, stream)
	}

	return
}

func (s *streamPool) BroadcastAsync(message *spacesyncproto.ObjectSyncMessage) (err error) {
	streams := s.getAllStreams()
	for _, stream := range streams {
		if err = stream.Send(message); err != nil {
			// TODO: add logging
		}
	}

	return nil
}

func (s *streamPool) AddAndReadStreamAsync(stream spacesyncproto.SpaceStream) {
	go s.AddAndReadStreamSync(stream)
}

func (s *streamPool) AddAndReadStreamSync(stream spacesyncproto.SpaceStream) (err error) {
	s.Lock()
	peerId, err := GetPeerIdFromStreamContext(stream.Context())
	if err != nil {
		s.Unlock()
		return
	}

	s.peerStreams[peerId] = stream
	s.wg.Add(1)
	s.Unlock()

	return s.readPeerLoop(peerId, stream)
}

func (s *streamPool) Close() (err error) {
	s.Lock()
	wg := s.wg
	s.Unlock()
	if wg != nil {
		wg.Wait()
	}
	return nil
}

func (s *streamPool) readPeerLoop(peerId string, stream spacesyncproto.SpaceStream) (err error) {
	defer s.wg.Done()
	limiter := make(chan struct{}, maxSimultaneousOperationsPerStream)
	for i := 0; i < maxSimultaneousOperationsPerStream; i++ {
		limiter <- struct{}{}
	}

	process := func(msg *spacesyncproto.ObjectSyncMessage) {
		if msg.TrackingId == "" {
			s.messageHandler(stream.Context(), peerId, msg)
			return
		}

		s.waitersMx.Lock()
		waiter, exists := s.waiters[msg.TrackingId]

		if !exists {
			s.waitersMx.Unlock()
			s.messageHandler(stream.Context(), peerId, msg)
			return
		}

		delete(s.waiters, msg.TrackingId)
		s.waitersMx.Unlock()
		waiter.ch <- msg
	}

Loop:
	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}
		select {
		case <-limiter:
		case <-stream.Context().Done():
			break Loop
		}
		go func() {
			process(msg)
			limiter <- struct{}{}
		}()
	}
	return s.removePeer(peerId)
}

func (s *streamPool) removePeer(peerId string) (err error) {
	s.Lock()
	defer s.Unlock()
	_, ok := s.peerStreams[peerId]
	if !ok {
		return ErrEmptyPeer
	}
	delete(s.peerStreams, peerId)
	return
}

func GetPeerIdFromStreamContext(ctx context.Context) (string, error) {
	conn, ok := ctx.Value(drpcctx.TransportKey{}).(sec.SecureConn)
	if !ok {
		return "", fmt.Errorf("incorrect connection type in stream")
	}

	return conn.RemotePeer().String(), nil
}

func genStreamPoolKey(peerId, treeId string, counter uint64) string {
	return fmt.Sprintf("%s.%s.%d", peerId, treeId, counter)
}