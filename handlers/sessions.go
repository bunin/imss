package handlers

import (
	"database/sql"
	"net/http"

	"github.com/bunin/imss/data"
	"github.com/bunin/imss/db"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.etcd.io/bbolt"
)

func ListScenes(ctx *gin.Context) {
	result := make([]*data.Session, 0, 20)
	if err := db.Get().View(func(tx *bbolt.Tx) error {
		c := tx.Bucket([]byte(db.BucketSessions)).Cursor()
		i := 0
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			i++
			if i >= 20 {
				break
			}
			s := &data.Session{}
			if err := proto.Unmarshal(v, s); err != nil {
				return err
			}
			images, err := loadImagesByScene(s.Id)
			if err != nil {
				return err
			}
			s.Images = images
			result = append(result, s)
		}
		return nil
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
	return
}

func GetScene(ctx *gin.Context) {
	sceneID := ctx.Param("id")
	scene, err := loadScene(sceneID, true)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Status(http.StatusNotFound)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	ctx.JSON(http.StatusOK, scene)
}

func CreateScene(ctx *gin.Context) {
	if s := data.GetActiveSession(); s != nil {
		ctx.Status(http.StatusFound)
		ctx.Header("Location", "/scene/"+s.Id)
		return
	}
	scene := &data.Session{
		Id:        xid.New().String(),
		Completed: false,
		CreatedAt: ptypes.TimestampNow(),
	}
	if err := db.Get().Update(func(tx *bbolt.Tx) error {
		b, err := proto.Marshal(scene)
		if err != nil {
			return err
		}
		return tx.Bucket([]byte(db.BucketSessions)).Put([]byte(scene.Id), b)
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	data.SetActiveSession(scene)
	ctx.Header("Location", "/scene/"+scene.Id)
	ctx.JSON(http.StatusCreated, scene)
}

func UpdateScene(ctx *gin.Context) {
	sceneID := ctx.Param("id")
	sceneToUpdate, err := loadScene(sceneID, false)
	newData := &data.Session{}
	if err := ctx.BindJSON(newData); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Status(http.StatusNotFound)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	currentScene := data.GetActiveSession()
	if !newData.Completed && currentScene != nil && currentScene.Id != sceneID && !currentScene.Completed {
		ctx.Header("Location", "/scene/"+currentScene.Id)
		ctx.Status(http.StatusFound)
		return
	}
	sceneToUpdate.Completed = newData.Completed
	if !newData.Completed {
		data.SetActiveSession(sceneToUpdate)
	} else {
		sceneToUpdate.FinishedAt = ptypes.TimestampNow()
		data.SetActiveSession(nil)
	}
	if err := sceneToUpdate.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
}

func loadScene(id string, withImages bool) (*data.Session, error) {
	s := &data.Session{}
	if err := db.Get().View(func(tx *bbolt.Tx) error {
		d := tx.Bucket([]byte(db.BucketSessions)).Get([]byte(id))
		if len(d) < 1 {
			return sql.ErrNoRows
		}
		if err := proto.Unmarshal(d, s); err != nil {
			return err
		}
		if withImages {
			images, err := loadImagesByScene(id)
			if err != nil {
				return err
			}
			s.Images = images
		}
		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "failed to load scene")
	}
	return s, nil
}
