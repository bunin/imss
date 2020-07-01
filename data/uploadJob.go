package data

import (
	"context"
	"net/url"
	"time"

	"github.com/bunin/imss/db"
	"github.com/bunin/imss/google"
	"github.com/bunin/imss/mail"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/gphotosuploader/googlemirror/api/photoslibrary/v1"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

func (x *UploadJob) Save() error {
	if x.Id == "" {
		x.Id = xid.New().String()
	}
	if x.CreatedAt == nil || !x.CreatedAt.IsValid() {
		x.CreatedAt = ptypes.TimestampNow()
	}
	xClone := proto.Clone(x).(*UploadJob)
	xClone.Images = nil
	return db.Get().Update(func(tx *bbolt.Tx) error {
		data, err := proto.Marshal(xClone)
		if err != nil {
			return errors.Wrap(err, "marshal upload job")
		}
		return errors.Wrap(
			tx.Bucket([]byte(db.BucketUploads)).Put([]byte(xClone.Id), data),
			"put upload job",
		)
	})
}

func (x *UploadJob) Run() {
	client, err := google.GetGoogleClient()
	l := zap.L()
	l.Info("RUN " + x.Id)
	if err != nil {
		x.Error = &Error{
			Code:    1,
			Message: err.Error(),
		}
		x.Status = UploadStatus_ERROR
		l.Error("get google client", zap.Error(err))
		if err := x.Save(); err != nil {
			l.Error("save upload error", zap.Error(err))
		}
		return
	}
	name := "Фото от " + time.Unix(x.CreatedAt.Seconds, int64(x.CreatedAt.Nanos)).Format("02.01.2006 15:04")
	if x.Name != "" {
		name = x.Name
	}
	album, err := client.GetOrCreateAlbumByName(name)
	if err != nil {
		l.Error("create album", zap.Error(err))
		x.Error = &Error{
			Code:    2,
			Message: err.Error(),
		}
		x.Status = UploadStatus_ERROR
		if err := x.Save(); err != nil {
			l.Error("save upload error", zap.Error(err))
		}
		return
	}
	for _, v := range x.Images {
		img := ImageByID(v.ImageId)
		if img == nil { // image not found
			v.Error = ErrImageUploadNoLocalImage
			v.Status = UploadStatus_ERROR
			if err := v.Save(); err != nil {
				l.Error("load image", zap.Error(err))
			}
			x.Status = UploadStatus_ERROR
			x.Error = ErrImageUploadNoLocalImage
			if err := x.Save(); err != nil {
				l.Error("load image", zap.Error(err))
			}
			return
		}
		// image found, invalid path
		localPath, err := url.PathUnescape(img.LocalPath)
		if err != nil {
			v.Error = ErrImageUploadNoLocalImage
			v.Error.Message = err.Error()
			v.Status = UploadStatus_ERROR
			if err := v.Save(); err != nil {
				l.Error("save image path", zap.Error(err))
			}
			x.Status = UploadStatus_ERROR
			x.Error = ErrImageUploadNoLocalImage
			x.Error.Message = err.Error()
			if err := x.Save(); err != nil {
				l.Error("save image path", zap.Error(err))
			}
			return
		}
		mi, err := client.AddMediaItem(context.Background(), localPath, album.Id)
		if err != nil {
			v.Error = ErrImageUploadNoLocalImage
			v.Error.Message = err.Error()
			v.Status = UploadStatus_ERROR
			if err := v.Save(); err != nil {
				l.Error("upload image", zap.Error(err))
			}
			x.Status = UploadStatus_ERROR
			x.Error = ErrImageUploadNoLocalImage
			x.Error.Message = err.Error()
			if err := x.Save(); err != nil {
				l.Error("upload image", zap.Error(err))
			}
			return
		}
		v.Status = UploadStatus_DONE
		v.Progress = img.Size
		x.Progress += img.Size
		v.CloudId = mi.Id
		if err := v.Save(); err != nil {
			l.Error("save image while upload", zap.Error(err))
		}
		if err := x.Save(); err != nil {
			l.Error("save image job while image upload", zap.Error(err))
		}
	}
	resp, err := client.Albums.Share(album.Id, &photoslibrary.ShareAlbumRequest{
		SharedAlbumOptions: &photoslibrary.SharedAlbumOptions{},
	}).Do()
	if err != nil {
		x.Status = UploadStatus_ERROR
		x.Error = &Error{
			Message: err.Error(),
		}
		if err := x.Save(); err != nil {
			l.Error("save image job while sharing an album", zap.Error(err))
		}
		return
	}
	x.CloudLink = resp.ShareInfo.ShareableUrl
	x.Status = UploadStatus_DONE
	x.FinishedAt = ptypes.TimestampNow()
	if err := x.Save(); err != nil {
		l.Error("save image job after upload", zap.Error(err))
	}
	if err := mail.Send(x.Recipient, name, x.CloudLink); err != nil {
		l.Error("send mail", zap.Error(err))
	}
}
