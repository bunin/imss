package handlers

import (
	"bytes"
	"net/http"

	"github.com/bunin/imss/data"
	"github.com/bunin/imss/db"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/gphotosuploader/google-photos-api-client-go/lib-gphotos"
	"github.com/gphotosuploader/googlemirror/api/photoslibrary/v1"
	"go.etcd.io/bbolt"
)

func Test(ctx *gin.Context) {
	hc := oc.Client(ctx, t)
	client, err := gphotos.NewClient(hc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	res, err := client.MediaItems.Search(&photoslibrary.SearchMediaItemsRequest{}).Do()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

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
