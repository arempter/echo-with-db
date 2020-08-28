package server

import (
	"github.com/sirupsen/logrus"
)

type C interface {
	GetLogLevel() logrus.Level
}

type Config struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"1234"`
	LogLevel string `env:"LOG_LEVEL"`
}

func (s Config) GetLogLevel() logrus.Level {
	switch s.LogLevel {
	case "DEBUG":
		return logrus.DebugLevel
	default:
		return logrus.InfoLevel
	}
}
