package data

import (
	"bytes"

	"github.com/bunin/imss/db"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

func (m *Image) Save() error {
	img, err := loadImageByPath(m.SessionId, m.LocalPath)
	if err != nil {
		return errors.Wrap(err, "loading image")
	}
	if img == nil { // new image
		img = m
		img.Id = img.SessionId + ":" + xid.New().String()
		img.CreatedAt = ptypes.TimestampNow()
		return img.saveToDB()
	}
	if img.Size == m.Size { // existing image with no changes
		return nil
	}
	img.Size = m.Size // existing image with changed size ??
	return img.saveToDB()
}

func (m *Image) saveToDB() error {
	return errors.Wrap(db.Get().Update(func(tx *bbolt.Tx) error {
		d, err := proto.Marshal(m)
		if err != nil {
			return errors.Wrap(err, "marshal image")
		}
		return errors.Wrap(tx.Bucket([]byte(db.BucketImages)).Put([]byte(m.Id), d), "put image")
	}), "save image")
}

func loadImageByPath(sessionID, path string) (*Image, error) {
	var result *Image
	if err := db.Get().View(func(tx *bbolt.Tx) error {
		c := tx.Bucket([]byte(db.BucketImages)).Cursor()
		prefix := []byte(sessionID + ":")
		img := &Image{}
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			if err := proto.Unmarshal(v, img); err != nil {
				return errors.Wrap(err, "unmarshal image")
			}
			if img.LocalPath != path {
				continue
			}
			result = img
			break
		}
		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "load image")
	}
	return result, nil
}

func ImageByID(id string) *Image {
	l := zap.L()
	var img *Image
	if err := db.Get().View(func(tx *bbolt.Tx) error {
		d := tx.Bucket([]byte(db.BucketImages)).Get([]byte(id))
		t := &Image{}
		if err := proto.Unmarshal(d, t); err != nil {
			return err
		}
		img = t
		return nil
	}); err != nil {
		l.Error("ImageByID", zap.Error(err))
	}
	return img
}
