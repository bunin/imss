package data

import (
	"sync"

	"github.com/bunin/imss/db"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/rs/xid"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

var (
	asLock        sync.RWMutex
	asOnce        sync.Once
	activeSession *Session
)

func GetActiveSession() *Session {
	asOnce.Do(func() {
		if err := LoadActiveSession(); err != nil {
			zap.L().Error("failed to load active session from db", zap.Error(err))
		}
	})
	asLock.RLock()
	defer asLock.RUnlock()
	return activeSession
}

func SetActiveSession(s *Session) {
	asLock.Lock()
	activeSession = s
	asLock.Unlock()
}

func LoadActiveSession() error {
	return db.Get().View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(db.BucketSessions)).ForEach(func(k, v []byte) error {
			s := &Session{}
			if err := proto.Unmarshal(v, s); err != nil {
				return err
			}
			if s.Active {
				SetActiveSession(s)
				return nil
			}
			return nil
		})
	})
}

func (m *Session) Save() error {
	if m.Id == "" { // new session
		m.Id = xid.New().String()
		m.CreatedAt = ptypes.TimestampNow()
	}

	return nil
}
