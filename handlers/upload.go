package handlers

import (
	"net/http"

	"github.com/bunin/imss/data"
	"github.com/bunin/imss/db"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.etcd.io/bbolt"
)

func Upload(ctx *gin.Context) {
	req := &data.UploadJob{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	if len(req.Images) < 1 {
		ctx.Status(http.StatusBadRequest)
		return
	}
	req.Status = data.UploadStatus_IN_PROGRESS
	req.Id = xid.New().String()
	for i, iu := range req.Images {
		if iu.Id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}
		img := &data.Image{}
		if err := db.Get().View(func(tx *bbolt.Tx) error {
			imgData := tx.Bucket([]byte(db.BucketImages)).Get([]byte(iu.Id))
			if len(imgData) < 1 {
				return errors.New("image " + iu.Id + " not found")
			}
			if err := proto.Unmarshal(imgData, img); err != nil {
				return errors.Wrap(err, "image "+iu.Id+" not found")
			}
			return nil
		}); err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{"error": errors.Wrap(err, "get image").Error()})
			return
		}
		req.Size += img.Size
		req.Images[i].ImageId = iu.Id
		req.Images[i].Id = req.Id + ":" + xid.New().String()
		req.Images[i].JobId = req.Id
		if err := req.Images[i].Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{"error": errors.Wrap(err, "save image").Error()})
			return
		}
		if err := req.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{"error": errors.Wrap(err, "save job").Error()})
			return
		}
	}
	go req.Run()
	ctx.JSON(http.StatusCreated, req)
}
