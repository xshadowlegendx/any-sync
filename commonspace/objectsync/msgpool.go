package objectsync

import (
	"context"
	"fmt"
	"github.com/anytypeio/any-sync/app/ocache"
	"github.com/anytypeio/any-sync/commonspace/objectsync/synchandler"
	"github.com/anytypeio/any-sync/commonspace/spacesyncproto"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type StreamManager interface {
	SendPeer(ctx context.Context, peerId string, msg *spacesyncproto.ObjectSyncMessage) (err error)
	SendResponsible(ctx context.Context, msg *spacesyncproto.ObjectSyncMessage) (err error)
	Broadcast(ctx context.Context, msg *spacesyncproto.ObjectSyncMessage) (err error)
}

// MessagePool can be made generic to work with different streams
type MessagePool interface {
	ocache.ObjectLastUsage
	synchandler.SyncHandler
	StreamManager
	SendSync(ctx context.Context, peerId string, message *spacesyncproto.ObjectSyncMessage) (reply *spacesyncproto.ObjectSyncMessage, err error)
}

type MessageHandler func(ctx context.Context, senderId string, message *spacesyncproto.ObjectSyncMessage) (err error)

type responseWaiter struct {
	ch chan *spacesyncproto.ObjectSyncMessage
}

type messagePool struct {
	sync.Mutex
	StreamManager
	messageHandler MessageHandler
	waiters        map[string]responseWaiter
	waitersMx      sync.Mutex
	counter        atomic.Uint64
	lastUsage      atomic.Int64
}

func newMessagePool(streamManager StreamManager, messageHandler MessageHandler) MessagePool {
	s := &messagePool{
		StreamManager:  streamManager,
		messageHandler: messageHandler,
		waiters:        make(map[string]responseWaiter),
	}
	return s
}

func (s *messagePool) SendSync(ctx context.Context, peerId string, msg *spacesyncproto.ObjectSyncMessage) (reply *spacesyncproto.ObjectSyncMessage, err error) {
	s.updateLastUsage()
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	newCounter := s.counter.Add(1)
	msg.ReplyId = genReplyKey(peerId, msg.ObjectId, newCounter)
	log.InfoCtx(ctx, "mpool sendSync", zap.String("replyId", msg.ReplyId))
	s.waitersMx.Lock()
	waiter := responseWaiter{
		ch: make(chan *spacesyncproto.ObjectSyncMessage, 1),
	}
	s.waiters[msg.ReplyId] = waiter
	s.waitersMx.Unlock()

	err = s.SendPeer(ctx, peerId, msg)
	if err != nil {
		return
	}
	select {
	case <-ctx.Done():
		s.waitersMx.Lock()
		delete(s.waiters, msg.ReplyId)
		s.waitersMx.Unlock()

		log.With(zap.String("replyId", msg.ReplyId)).WarnCtx(ctx, "time elapsed when waiting")
		err = fmt.Errorf("sendSync context error: %v", ctx.Err())
	case reply = <-waiter.ch:
		// success
	}
	return
}

func (s *messagePool) SendPeer(ctx context.Context, peerId string, msg *spacesyncproto.ObjectSyncMessage) (err error) {
	s.updateLastUsage()
	return s.StreamManager.SendPeer(ctx, peerId, msg)
}

func (s *messagePool) SendResponsible(ctx context.Context, msg *spacesyncproto.ObjectSyncMessage) (err error) {
	s.updateLastUsage()
	return s.StreamManager.SendResponsible(ctx, msg)
}
func (s *messagePool) Broadcast(ctx context.Context, msg *spacesyncproto.ObjectSyncMessage) (err error) {
	s.updateLastUsage()
	return s.StreamManager.Broadcast(ctx, msg)
}

func (s *messagePool) HandleMessage(ctx context.Context, senderId string, msg *spacesyncproto.ObjectSyncMessage) (err error) {
	s.updateLastUsage()
	if msg.ReplyId != "" {
		log.InfoCtx(ctx, "mpool receive reply", zap.String("replyId", msg.ReplyId))
		// we got reply, send it to waiter
		if s.stopWaiter(msg) {
			return
		}
		log.DebugCtx(ctx, "reply id does not exist", zap.String("replyId", msg.ReplyId))
	}
	return s.messageHandler(ctx, senderId, msg)
}

func (s *messagePool) LastUsage() time.Time {
	return time.Unix(s.lastUsage.Load(), 0)
}

func (s *messagePool) updateLastUsage() {
	s.lastUsage.Store(time.Now().Unix())
}

func (s *messagePool) stopWaiter(msg *spacesyncproto.ObjectSyncMessage) bool {
	s.waitersMx.Lock()
	waiter, exists := s.waiters[msg.ReplyId]
	if exists {
		delete(s.waiters, msg.ReplyId)
		s.waitersMx.Unlock()
		waiter.ch <- msg
		return true
	}
	s.waitersMx.Unlock()
	return false
}

func genReplyKey(peerId, treeId string, counter uint64) string {
	b := &strings.Builder{}
	b.WriteString(peerId)
	b.WriteString(".")
	b.WriteString(treeId)
	b.WriteString(".")
	b.WriteString(strconv.FormatUint(counter, 36))
	return b.String()
}