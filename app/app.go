package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bunin/imss/data"
	"github.com/bunin/imss/db"
	"github.com/bunin/imss/handlers"
	"github.com/fsnotify/fsnotify"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type app struct {
	cfg Config
	w   *fsnotify.Watcher
	srv *http.Server
}

func New(cfg Config) *app {
	return &app{cfg: cfg}
}

func (a *app) Run() error {
	var (
		err    error
		logger *zap.Logger
	)
	if logger, err = a.loggerConfig().Build(); err != nil {
		return errors.Wrap(err, "failed to init logger")
	}
	zap.ReplaceGlobals(logger)
	if err = db.Init(a.cfg.DB); err != nil {
		return err
	}
	if a.w, err = fsnotify.NewWatcher(); err != nil {
		return errors.Wrap(err, "failed to create a file watcher")
	}
	if err = a.w.Add(a.cfg.Dir); err != nil {
		return errors.Wrap(err, "failed to watch "+a.cfg.Dir+" dir")
	}

	if err = a.startServer(); err != nil {
		return errors.Wrap(err, "failed to start a server")
	}

	a.watch()

	return nil // todo
}

func (a *app) Stop() error {
	var e error
	if err := a.w.Close(); err != nil {
		e = multierror.Append(e, errors.Wrap(err, "failed to close file watcher"))
	}
	if err := db.Close(); err != nil {
		e = multierror.Append(e, errors.Wrap(err, "failed to close database"))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.srv.Shutdown(ctx); err != nil {
		e = multierror.Append(e, errors.Wrap(err, "failed to stop http server"))
	}
	return e
}

func (a *app) watch() {
	logger := zap.L()
	for {
		select {
		case err := <-a.w.Errors:
			{
				if err != nil {
					logger.Error("watch error", zap.Error(err))
					return
				}
			}
		case event, more := <-a.w.Events:
			if !more {
				logger.Info("watcher closed")
				return
			}
			switch event.Op.String() {
			case "WRITE":
				f, err := os.Stat(event.Name)
				if err != nil {
					log.Println("failed to open file", err)
					continue
				}
				if f.IsDir() {
					logger.Info(event.Name + " is a directory, skipping")
					continue
				}
				scene := data.GetActiveSession()
				img := &data.Image{
					SessionId: scene.Id,
					LocalPath: event.Name,
					Size:      uint64(f.Size()),
				}
				if err := img.Save(); err != nil {
					logger.Error("failed to save local image", zap.Error(err))
				}
			}
		}
	}
}

func (a *app) loggerConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(a.cfg.LogLevel),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{a.cfg.LogFile},
		ErrorOutputPaths: []string{a.cfg.LogFile},
	}
}

func (a *app) startServer() error {
	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false), ginzap.RecoveryWithZap(zap.L(), true))

	r.GET("/session", handlers.ListScenes)
	r.GET("/session/:id", handlers.GetScene)
	r.PATCH("/session/:id", handlers.UpdateScene)
	r.POST("/session", handlers.CreateScene)
	r.POST("/upload", handlers.Upload)

	a.srv = &http.Server{Addr: ":" + strconv.FormatUint(a.cfg.Port, 10), Handler: r}

	go func() {
		if err := a.srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()
	return nil
}
