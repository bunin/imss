package app

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bunin/imss/data/config"
	"github.com/bunin/imss/db"
	"github.com/bunin/imss/google"
	"github.com/bunin/imss/handlers"
	"github.com/bunin/imss/mail"
	"github.com/bunin/imss/ui"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type app struct {
	cfg *Config
	srv *http.Server
}

func New(cfg *Config) *app {
	return &app{cfg: cfg}
}

func (a *app) Run() error {
	var (
		err    error
		logger *zap.Logger
	)
	if logger, err = a.loggerConfig().Build(); err != nil {
		return errors.Wrap(err, "init logger")
	}
	zap.ReplaceGlobals(logger)
	if err = db.Init(a.cfg.DB); err != nil {
		return err
	}
	if err = a.startServer(); err != nil {
		return errors.Wrap(err, "starting a server")
	}
	config.SetPhoneDir(a.cfg.PhoneDir)
	config.SetTargetDir(a.cfg.LocalDir)
	mail.Setup(
		a.cfg.SMTPHost,
		a.cfg.SMTPPort,
		a.cfg.SMTPFromName,
		a.cfg.SMTPFromEmail,
		a.cfg.SMTPLogin,
		a.cfg.SMTPPassword,
	)
	return nil
}

func (a *app) Stop() error {
	var e error
	if err := db.Close(); err != nil {
		e = multierror.Append(e, errors.Wrap(err, "closing database"))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.srv.Shutdown(ctx); err != nil {
		e = multierror.Append(e, errors.Wrap(err, "stopping http server"))
	}
	return e
}

func (a *app) loggerConfig() zap.Config {
	return zap.Config{
		Level:             zap.NewAtomicLevelAt(a.cfg.LogLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "console",
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       []string{a.cfg.LogFile},
		ErrorOutputPaths:  []string{a.cfg.LogFile},
	}
}

func (a *app) startServer() error {
	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false), ginzap.RecoveryWithZap(zap.L(), true))
	r.NoRoute(ui.Handle)
	r.GET("/auth", google.Auth)
	r.GET("/files/*path", handlers.ServeImage)
	r.HEAD("/files/*path", handlers.ServeImage)
	google.SetID(a.cfg.GoogleClientID)
	google.SetSecret(a.cfg.GoogleSecret)
	api := r.Group("/api")
	{
		api.GET("/auth/check", google.CheckAuth(a.cfg.GoogleClientID, a.cfg.GoogleSecret))
		api.GET("/session", handlers.ListSessions)
		api.GET("/session/:id", handlers.GetSession)
		api.PATCH("/session/:id", handlers.UpdateSession)
		api.POST("/session", handlers.CreateSession)
		api.POST("/upload", handlers.Upload)
	}

	a.srv = &http.Server{Addr: ":" + strconv.FormatUint(a.cfg.Port, 10), Handler: r}

	go func() {
		if err := a.srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()
	return nil
}
