package data

import (
	"github.com/bunin/imss/db"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

var (
	ErrImageUploadNoLocalImage = &Error{Code: 101, Message: "Image to upload not found"}
)

func (x *ImageUpload) Save() error {
	return db.Get().Update(func(tx *bbolt.Tx) error {
		data, err := proto.Marshal(x)
		if err != nil {
			return errors.Wrap(err, "marshal image upload")
		}
		return errors.Wrap(
			tx.Bucket([]byte(db.BucketImageUploads)).Put([]byte(x.Id), data),
			"put image upload",
		)
	})
}
