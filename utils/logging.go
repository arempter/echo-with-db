package utils

import (
	"echo-with-db/errors"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

func Log(e error) {
	err, ok := e.(*errors.Error)
	if !ok {
		log.Error(err)
		return
	}

	log := logrus.WithFields(logrus.Fields{
		"operation": err.Op,
	})

	logrus.SetFormatter(&logrus.JSONFormatter{})

	switch err.Level {
	case logrus.ErrorLevel:
		log.Error(err.Msg)
	case logrus.DebugLevel:
		log.Debug(err.Msg)
	case logrus.InfoLevel:
		log.Info(err.Msg)
	default:
		log.Info(err.Msg)
	}
}
