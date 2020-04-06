package handlers

import (
	"bytes"

	"github.com/bunin/imss/data"
	"github.com/bunin/imss/db"
	"github.com/golang/protobuf/proto"
	"go.etcd.io/bbolt"
)

func loadImagesByScene(sceneID string) ([]*data.Image, error) {
	images := make([]*data.Image, 0, 16)
	if err := db.Get().View(func(tx *bbolt.Tx) error {
		imagesCursor := tx.Bucket([]byte(db.BucketImages)).Cursor()
		prefix := []byte(sceneID)
		for ik, iv := imagesCursor.Seek(prefix); ik != nil && bytes.HasPrefix(ik, prefix); ik, iv = imagesCursor.Next() {
			i := &data.Image{}
			if err := proto.Unmarshal(iv, i); err != nil {
				return err
			}
			images = append(images, i)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return images, nil
}
