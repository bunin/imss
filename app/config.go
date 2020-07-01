package app

import (
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	DB             string        `env:"IMSS_DB" envDefault:"imss.db"`
	GoogleClientID string        `env:"IMSS_GOOGLE_ID"`
	GoogleSecret   string        `env:"IMSS_GOOGLE_SECRET"`
	LocalDir       string        `env:"LOCAL_DIR" envDefault:"C:\\tmp"`
	LogFile        string        `env:"IMSS_LOG" envDefault:"imss.log"`
	LogLevel       zapcore.Level `env:"IMSS_LOG_LEVEL" envDefault:"debug"`
	PhoneDir       string        `env:"PHONE_DIR" envDefault:"/sdcard/DCIM/Camera/"`
	Port           uint64        `env:"IMSS_PORT" envDefault:"8080"`
	SMTPFromEmail  string        `env:"SMTP_FROM_EMAIL"`
	SMTPFromName   string        `env:"SMTP_FROM_NAME"`
	SMTPHost       string        `env:"SMTP_HOST" envDefault:"smtp.gmail.com"`
	SMTPLogin      string        `env:"SMTP_LOGIN"`
	SMTPPassword   string        `env:"SMTP_PASSWORD"`
	SMTPPort       int           `env:"SMTP_PORT" envDefault:"587"`
}

func (c *Config) Validate() error {
	if c.DB == "" {
		return errors.New("IMSS_DB is required")
	}
	if c.GoogleClientID == "" {
		return errors.New("IMSS_GOOGLE_ID is required")
	}
	if c.GoogleSecret == "" {
		return errors.New("IMSS_GOOGLE_SECRET is required")
	}
	if c.SMTPLogin == "" {
		return errors.New("SMTP_LOGIN is required")
	}
	if c.SMTPPassword == "" {
		return errors.New("SMTP_PASSWORD is required")
	}
	if c.SMTPFromEmail == "" {
		return errors.New("SMTP_FROM_EMAIL is required")
	}
	if c.SMTPFromEmail == "" {
		return errors.New("SMTP_FROM_NAME is required")
	}
	return nil
}
