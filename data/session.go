package data

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/bunin/imss/data/config"
	"github.com/bunin/imss/db"
	"github.com/bunin/imss/photos"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
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
			zap.L().Error("load active session from db", zap.Error(err))
		}
	})
	asLock.RLock()
	defer asLock.RUnlock()
	return activeSession
}

func SetActiveSession(s *Session) {
	asLock.Lock()
	activeSession = s
	go func() {
		if err := pullFiles(); err != nil {
			zap.L().Error("pull files", zap.Error(err))
		}
	}()
	asLock.Unlock()
}

func LoadActiveSession() error {
	return db.Get().View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(db.BucketSessions)).ForEach(func(k, v []byte) error {
			s := &Session{}
			if err := proto.Unmarshal(v, s); err != nil {
				return err
			}
			if s.GetIsActive() {
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

	return m.saveToDB()
}

func (m *Session) saveToDB() error {
	if m == nil {
		return errors.New("saving a nil session")
	}
	return db.Get().Update(func(tx *bbolt.Tx) error {
		sCopy := proto.Clone(m).(*Session)
		sCopy.Images = nil
		b, err := proto.Marshal(sCopy)
		if err != nil {
			return err
		}
		return errors.Wrap(tx.Bucket([]byte(db.BucketSessions)).Put([]byte(sCopy.Id), b), "save session")
	})
}

func (m *Session) addImage(i *Image) error {
	m.Images = append(m.Images, i)
	i.SessionId = m.Id
	return i.Save()
}

func pullFiles() error {
	l := zap.L()
	as := GetActiveSession()
	if as == nil {
		l.Info("no active session - nothing to pull")
		return nil
	}
	data, err := photos.Listen(config.GetPhoneDir())
	if err != nil {
		l.Error("Listen", zap.Error(err))
	}
	for path := range data {
		target := filepath.Join(config.GetTargetDir(), filepath.Base(path))
		cmd := exec.Command("adb", "pull", path, target)
		stdErr, stdOut := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
		cmd.Stderr = stdErr
		cmd.Stdout = stdOut
		if err := cmd.Run(); err != nil {
			l.Error(cmd.String(), zap.Error(err))
			continue
		}
		if stdOut.Len() > 0 {
			l.Info(cmd.String() + ": " + stdOut.String())
		}
		if stdErr.Len() > 0 {
			l.Error(cmd.String() + ": " + stdErr.String())
		}
		stat, err := os.Stat(target)
		if err != nil {
			l.Error("stat file", zap.Error(err))
			continue
		}
		if stat.Size() < 1048576 { // do not save files less than 1MB
			continue
		}
		img := &Image{
			Id:        xid.New().String(),
			SessionId: as.Id,
			LocalPath: target,
			CreatedAt: ptypes.TimestampNow(),
			Size:      uint64(stat.Size()),
		}
		if err := as.addImage(img); err != nil {
			l.Error("add image to session", zap.Error(err), zap.String("sessID", as.Id))
			continue
		}
	}
	return nil
}
