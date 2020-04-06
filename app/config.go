package app

import (
	"os"

	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	AccessKey  string        `env:"IMSS_KEY"`
	BucketName string        `env:"IMSS_BUCKET" envDefault:"pics"`
	DB         string        `env:"IMSS_DB" envDefault:"imss.db"`
	Dir        string        `env:"IMSS_DIR" envDefault:"."`
	LogFile    string        `env:"IMSS_LOG" envDefault:"imss.log"`
	LogLevel   zapcore.Level `env:"IMSS_LOG_LEVEL" envDefault:"debug"`
	Port       uint64        `env:"IMSS_PORT" envDefault:"8080"`
	SecretKey  string        `env:"IMSS_SECRET"`
}

func (c Config) Validate() error {
	if c.DB == "" {
		return errors.New("IMSS_DB is required")
	}
	if c.Dir == "" {
		return errors.New("IMSS_DIR is required")
	}
	fi, err := os.Stat(c.Dir)
	if err != nil {
		return errors.Wrap(err, "failed to stat "+c.Dir)
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
