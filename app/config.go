package app

import (
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	Dir        string `env:"S3P_DIR" envDefault:"."`
	AccessKey  string `env:"S3P_KEY"`
	SecretKey  string `env:"S3P_SECRET"`
	BucketName string `env:"S3P_BUCKET" envDefault:"pics"`
}

func (c Config) Validate() error {
	if c.Dir == "" {
		return errors.New("S3P_DIR is required")
	}
	fi, err := os.Stat(c.Dir)
	if err != nil {
		return errors.Wrap(err, "Failed to stat "+c.Dir)
	}
	if !fi.IsDir() {
		return errors.New(c.Dir + " is not a dir")
	}
	if c.AccessKey == "" {
		return errors.New("S3P_KEY is required")
	}
	if c.SecretKey == "" {
		return errors.New("S3P_SECRET is required")
	}
	if c.BucketName == "" {
		return errors.New("S3P_BUCKET is required")
	}
	return nil
}
