package server

import (
	"echo-with-db/api"
	"echo-with-db/errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UsersService interface {
	Start()
	setupRoutes()
	attachMW()
}

type service struct {
	e  *echo.Echo
	mw []echo.MiddlewareFunc
	c  Config
}

func New(cfg Config, mw []echo.MiddlewareFunc) *service {
	return &service{
		e:  echo.New(),
		mw: mw,
		c:  cfg,
	}
}

func (s service) Start() *errors.Error {
	const op errors.Op = "service.start"
	const kind errors.Msg = "failed to start service"

	s.setupRoutes()
	s.attachMW()

	err := s.e.Start(fmt.Sprintf("%s:%d", s.c.Host, s.c.Port))
	if err != nil {
		return errors.E(op, kind, err, logrus.ErrorLevel)
	}
	return nil
}

func (s service) attachMW() {
	for _, f := range s.mw {
		s.e.Use(f)
	}
}

func (s service) setupRoutes() {
	s.e.Add("POST", "/user/add", api.HandlerAddUser())
}
