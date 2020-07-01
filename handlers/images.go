package handlers

import (
	"bytes"
	"net/url"
	"strings"

	"github.com/bunin/imss/data"
	"github.com/bunin/imss/db"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"go.etcd.io/bbolt"
)

func loadImagesBySession(sessionID string) ([]*data.Image, error) {
	images := make([]*data.Image, 0, 16)
	if err := db.Get().View(func(tx *bbolt.Tx) error {
		imagesCursor := tx.Bucket([]byte(db.BucketImages)).Cursor()
		prefix := []byte(sessionID)
		for ik, iv := imagesCursor.Seek(prefix); ik != nil && bytes.HasPrefix(ik, prefix); ik, iv = imagesCursor.Next() {
			i := &data.Image{}
			if err := proto.Unmarshal(iv, i); err != nil {
				return err
			}
			i.LocalPath = url.PathEscape(i.LocalPath)
			images = append(images, i)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return images, nil
}

func ServeImage(c *gin.Context) {
	path := strings.TrimPrefix(c.Param("path"), "/")
	c.File(path)
}
