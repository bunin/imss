package app

import (
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	Dir        string `env:"IMSS_DIR" envDefault:"."`
	AccessKey  string `env:"IMSS_KEY"`
	SecretKey  string `env:"IMSS_SECRET"`
	BucketName string `env:"IMSS_BUCKET" envDefault:"pics"`
}

func (c Config) Validate() error {
	if c.Dir == "" {
		return errors.New("IMSS_DIR is required")
	}
	fi, err := os.Stat(c.Dir)
	if err != nil {
		return errors.Wrap(err, "Failed to stat "+c.Dir)
	}
	if !fi.IsDir() {
		return errors.New(c.Dir + " is not a dir")
	}
	if c.AccessKey == "" {
		return errors.New("IMSS_KEY is required")
	}
	if c.SecretKey == "" {
		return errors.New("IMSS_SECRET is required")
	}
	if c.BucketName == "" {
		return errors.New("IMSS_BUCKET is required")
	}
	return nil
}
