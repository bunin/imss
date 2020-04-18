package app

import (
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	ClientID string        `env:"IMSS_ID"`
	DB       string        `env:"IMSS_DB" envDefault:"imss.db"`
	LogFile  string        `env:"IMSS_LOG" envDefault:"imss.log"`
	LogLevel zapcore.Level `env:"IMSS_LOG_LEVEL" envDefault:"debug"`
	Port     uint64        `env:"IMSS_PORT" envDefault:"8080"`
	Secret   string        `env:"IMSS_SECRET"`
}

func (c *Config) Validate() error {
	if c.DB == "" {
		return errors.New("IMSS_DB is required")
	}
	if c.ClientID == "" {
		return errors.New("IMSS_ID is required")
	}
	if c.Secret == "" {
		return errors.New("IMSS_SECRET is required")
	}
	return nil
}
