package server

import (
	"echo-with-db/api"
	"echo-with-db/config"
	"echo-with-db/errors"
	"echo-with-db/utils"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UsersService interface {
	SetupRoutes()
	Start()
	AttachMW()
}

type service struct {
	e  *echo.Echo
	mw []echo.MiddlewareFunc
	c  config.Server
}

func New(mw []echo.MiddlewareFunc) *service {
	cfg := config.Server{}
	if err := env.Parse(&cfg); err != nil {
		utils.Log(errors.E("failed to parse server configuration"))
	}
	return &service{
		e:  echo.New(),
		mw: mw,
		c:  cfg,
	}
}

func (s service) Start() *errors.Error {
	const op errors.Op = "service.start"
	const kind errors.Message = "failed to start service"

	s.SetupRoutes()
	s.AttachMW()
	err := s.e.Start(fmt.Sprintf("%s:%d", s.c.Host, s.c.Port))
	if err != nil {
		return errors.E(op, kind, err, logrus.ErrorLevel)
	}
	return nil
}

func (s service) AttachMW() {
	for _, f := range s.mw {
		s.e.Use(f)
	}
}

func (s service) SetupRoutes() {
	s.e.Add("POST", "/user/add", api.HandlerAddUser())
}
