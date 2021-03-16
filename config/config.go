package config

import (
	"github.com/sirupsen/logrus"
)

type C interface {
	GetLogLevel() logrus.Level
	GetDBDriver() string
}

type Config struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"1234"`
	LogLevel string `env:"LOG_LEVEL"`
	DbDriver string `env:"DB_DRIVER" envDefault:"mysql"`
}

func (s Config) GetDBDriver() string {
	return s.DbDriver
}

func (s Config) GetLogLevel() logrus.Level {
	switch s.LogLevel {
	case "DEBUG":
		return logrus.DebugLevel
	default:
		return logrus.InfoLevel
	}
}
