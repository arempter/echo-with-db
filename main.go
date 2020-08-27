package main

import (
	"echo-with-db/database"
	"echo-with-db/errors"
	customMW "echo-with-db/middleware"
	"echo-with-db/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"echo-with-db/utils"
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
	const msg errors.Message = "starting rest server"
	utils.Log(errors.E(op, msg))

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
	s := server.New(mw)

	err = s.Start()
	if err != nil {
		return err
	}

	return nil
}
