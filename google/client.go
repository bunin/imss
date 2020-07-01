package google

import (
	"context"
	"net/http"
	"strings"
	"sync"

	"github.com/bunin/imss/db"
	"github.com/bunin/imss/photos"
	"github.com/gin-gonic/gin"
	gphotos "github.com/gphotosuploader/google-photos-api-client-go/lib-gphotos"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

var (
	oc = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint:     endpoints.Google,
		RedirectURL:  "http://localhost:8080/auth", // todo
		Scopes:       photos.Scopes,
	}
	tOnce        sync.Once
	t            *oauth2.Token
	state        string
	googleClient *gphotos.Client
)

func SetID(s string) {
	oc.ClientID = s
}

func SetSecret(s string) {
	oc.ClientSecret = s
}

func CheckAuth(id, secret string) gin.HandlerFunc {
	oc.ClientID = id
	oc.ClientSecret = secret
	// authenticated already
	tOnce.Do(func() {
		if err := db.Get().View(func(tx *bbolt.Tx) error {
			data := tx.Bucket([]byte(db.BucketMisc)).Get([]byte("token"))
			if len(data) < 1 {
				return nil
			}
			t = &oauth2.Token{}
			return errors.Wrap(jsoniter.Unmarshal(data, t), "parse saved token")
		}); err != nil {
			zap.L().Error("load token", zap.Error(err))
		}
	})
	if t != nil {
		return func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, jsoniter.RawMessage(`{}`))
		}
	}
	return func(ctx *gin.Context) {
		// authenticated already
		if t != nil {
			ctx.Status(http.StatusOK)
			ctx.JSON(http.StatusOK, jsoniter.RawMessage(`{}`))
			return
		}
		state = xid.New().String()
		url := oc.AuthCodeURL(state, oauth2.AccessTypeOffline)
		ctx.JSON(http.StatusOK, map[string]string{"url": url})
	}
}

func Auth(ctx *gin.Context) {
	if ctx.Query("state") != state {
		state = "f"
		ctx.JSON(http.StatusUnauthorized, "auth failed")
		return
	}
	state = "q"
	scopes := strings.Split(ctx.Query("scope"), " ")
	for _, s := range photos.Scopes {
		found := false
		for _, s2 := range scopes {
			if s2 == s {
				found = true
				break
			}
		}
		if !found {
			ctx.JSON(http.StatusUnauthorized, "insufficient scopes")
			return
		}
	}
	var err error
	if t, err = oc.Exchange(context.Background(), ctx.Query("code")); err != nil {
		ctx.JSON(http.StatusUnauthorized, "invalid code")
		return
	} else {
		if err := db.Get().Update(func(tx *bbolt.Tx) error {
			data, e := jsoniter.Marshal(t)
			if e != nil {
				return errors.Wrap(e, "marshal token")
			}
			return errors.Wrap(tx.Bucket([]byte(db.BucketMisc)).Put([]byte("token"), data), "save token")
		}); err != nil {
			ctx.JSON(http.StatusInternalServerError, "authorize")
			return
		}
	}
	ctx.Redirect(http.StatusFound, "/")
	return
}

func GetGoogleClient() (*gphotos.Client, error) {
	if googleClient != nil {
		return googleClient, nil
	}
	var err error
	googleClient, err = gphotos.NewClient(oc.Client(context.Background(), t))
	return googleClient, err
}
