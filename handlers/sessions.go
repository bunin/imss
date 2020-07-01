package handlers

import (
	"database/sql"
	"net/http"

	"github.com/bunin/imss/data"
	"github.com/bunin/imss/db"
	"github.com/bunin/imss/photos"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

func ListSessions(ctx *gin.Context) {
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
			images, err := loadImagesBySession(s.Id)
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

func GetSession(ctx *gin.Context) {
	sessionID := ctx.Param("id")
	session, err := loadSession(sessionID, true)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Status(http.StatusNotFound)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	ctx.JSON(http.StatusOK, session)
}

func CreateSession(ctx *gin.Context) {
	if s := data.GetActiveSession(); s != nil {
		ctx.Status(http.StatusFound)
		ctx.Header("Location", "/api/session/"+s.Id)
		return
	}
	session := &data.Session{}
	if err := ctx.BindJSON(session); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	session.Id = xid.New().String()
	session.IsActive = true
	session.CreatedAt = ptypes.TimestampNow()
	session.Images = nil

	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	data.SetActiveSession(session)
	ctx.Header("Location", "/api/session/"+session.Id)
	ctx.JSON(http.StatusCreated, session)
}

func UpdateSession(ctx *gin.Context) {
	sessionID := ctx.Param("id")
	sessionToUpdate, err := loadSession(sessionID, false)
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
	currentSession := data.GetActiveSession()
	if newData.IsActive && currentSession != nil && currentSession.Id != sessionID && currentSession.IsActive {
		ctx.Header("Location", "/api/session/"+currentSession.Id)
		ctx.Status(http.StatusFound)
		return
	}
	sessionToUpdate.IsActive = newData.IsActive
	if newData.IsActive {
		data.SetActiveSession(sessionToUpdate)
	} else {
		sessionToUpdate.FinishedAt = ptypes.TimestampNow()
		if err := photos.Stop(); err != nil {
			zap.L().Error("stop listening", zap.Error(err))
		}
		data.SetActiveSession(nil)
	}
	if err := sessionToUpdate.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, sessionToUpdate)
}

func loadSession(id string, withImages bool) (*data.Session, error) {
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
			images, err := loadImagesBySession(id)
			if err != nil {
				return err
			}
			s.Images = images
		}
		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "load session")
	}
	return s, nil
}
