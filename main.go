package main

import (
	"echo-with-db/config"
	"echo-with-db/errors"
	"echo-with-db/server"
	"echo-with-db/utils"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := run(); err != nil {
		utils.Log(err)
		os.Exit(1)
	}
}

func run() *errors.Error {
	const op errors.Op = "server.main"

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.E(op, errors.Msg("failed to parse server configuration"))
	}

	logrus.SetLevel(cfg.GetLogLevel())
	utils.Log(errors.E(op, errors.Msg("starting rest server"), logrus.DebugLevel))

	s, err := server.New(cfg)
	if err != nil {
		return err
	}

	err = s.Start()
	if err != nil {
		return err
	}
	return nil
}
