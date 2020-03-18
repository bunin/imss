package app

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
)

type app struct {
	cfg      Config
	w        *fsnotify.Watcher
	uploader *s3manager.Uploader
}

func New(cfg Config) *app {
	return &app{cfg: cfg}
}

func (a *app) Run() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return errors.Wrap(err, "failed to create a file watcher")
	}
	if err = watcher.Add(a.cfg.Dir); err != nil {
		return errors.Wrap(err, "failed to watch "+a.cfg.Dir+" dir")
	}
	a.w = watcher

	sess := session.Must(session.NewSession(
		&aws.Config{Credentials: credentials.NewStaticCredentials(a.cfg.AccessKey, a.cfg.SecretKey, "")},
	))
	a.uploader = s3manager.NewUploader(sess)

	a.watch()

	return nil // todo
}

func (a *app) Stop() error {
	if err := a.w.Close(); err != nil {
		return errors.Wrap(err, "failed to close file watcher")
	}
	return nil
}

func (a *app) watch() {
	for {
		select {
		case err := <-a.w.Errors:
			{
				if err != nil {
					log.Println("watch error", err)
					return
				}
			}
		case event, more := <-a.w.Events:
			if !more {
				log.Println("watcher closed")
				return
			}
			switch event.Op.String() {
			case "WRITE":
				f, err := os.Open(event.Name)
				if err != nil {
					log.Println("failed to open file", err)
					continue
				}
				if _, err := a.uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(a.cfg.BucketName),
					Key:    aws.String(event.Name),
					Body:   f,
				}); err != nil {
					log.Println("failed to upload file "+event.Name, err)
				}
				if err := f.Close(); err != nil {
					log.Println("failed to close "+f.Name(), err)
				}
			}
		}
	}
}
