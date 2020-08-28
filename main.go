package main

import (
	"echo-with-db/database"
	"echo-with-db/errors"
	customMW "echo-with-db/middleware"
	"echo-with-db/server"
	"echo-with-db/utils"
	"github.com/caarlos0/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		utils.Log(err)
		os.Exit(1)
	}
}

func run() *errors.Error {
	const op errors.Op = "server.startup"

	cfg := server.Config{}
	if err := env.Parse(&cfg); err != nil {
		utils.Log(errors.E("failed to parse server configuration"))
	}

	logrus.SetLevel(cfg.GetLogLevel())
	utils.Log(errors.E(op, errors.Msg("starting rest server"), logrus.DebugLevel))

	db, err := database.New(true)
	if err != nil {
		return err
	}

	mw := []echo.MiddlewareFunc{
		customMW.AttachDatabase(db),
		customMW.ContextTimeout(3 * time.Second),
		middleware.RequestID(),
		middleware.Logger(),
	}
	s := server.New(cfg, mw)

	err = s.Start()
	if err != nil {
		return err
	}

	return nil
}
