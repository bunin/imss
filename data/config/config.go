package config

import (
	"log"
	"os"

	"go.uber.org/zap"
)

var (
	phoneDir  string
	targetDir string
)

func SetPhoneDir(s string) {
	phoneDir = s
}

func SetTargetDir(s string) {
	targetDir = s
	if err := os.MkdirAll(s, 0777); err != nil {
		log.Fatal("create target dir", zap.Error(err))
	}
}

func GetPhoneDir() string {
	return phoneDir
}

func GetTargetDir() string {
	return targetDir
}
