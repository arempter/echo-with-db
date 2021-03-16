package server

import (
	"echo-with-db/api"
	"echo-with-db/config"
	"echo-with-db/database"
	"echo-with-db/database/mysql"
	"echo-with-db/database/postgres"
	"echo-with-db/errors"
	customMW "echo-with-db/middleware"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

type Server interface {
	Start()
	setupRoutes()
	attachMW()
}

type server struct {
	e  *echo.Echo
	mw []echo.MiddlewareFunc
	c  config.Config
}

const MYSQL = "mysql"
const POSTGRES = "postgres"

func New(cfg config.Config) (*server, *errors.Error) {
	dbLayer, err := connect(cfg)
	if err != nil {
		return nil, err
	}
	mw := []echo.MiddlewareFunc{
		customMW.AttachDatabase(dbLayer),
		customMW.ContextTimeout(3 * time.Second),
		middleware.RequestID(),
		middleware.Logger(),
	}

	return &server{
		e:  echo.New(),
		mw: mw,
		c:  cfg,
	}, nil
}

func connect(cfg config.Config) (database.Database, *errors.Error) {
	const op errors.Op = "database.connect"
	switch cfg.DbDriver {
	case MYSQL:
		con, err := sqlx.Connect(MYSQL, "echo:echo@/echodb")
		if err != nil {
			return nil, errors.E(op, errors.Msg("failed to connect to database: "+err.Error()), logrus.ErrorLevel)
		}
		dbLayer, err := mysql.New(con, true, cfg)
		if err.(*errors.Error) != nil {
			return nil, err.(*errors.Error)
		}
		return dbLayer, nil
	case POSTGRES:
		con, err := sqlx.Connect(POSTGRES, "postgres://echo:echo@localhost:5432/echodb?sslmode=disable")
		if err != nil {
			return nil, errors.E(op, errors.Msg("failed to connect to database: "+err.Error()), logrus.ErrorLevel)
		}
		dbLayer, err := postgres.New(con, true, cfg)
		if err.(*errors.Error) != nil {
			return nil, err.(*errors.Error)
		}
		return dbLayer, nil
	default:
		return nil, errors.E(op, errors.Msg("unsupported database type"), logrus.ErrorLevel)
	}
}

func (s server) Start() *errors.Error {
	const op errors.Op = "server.start"
	const kind errors.Msg = "failed to start server"

	s.setupRoutes()
	s.attachMW()

	err := s.e.Start(fmt.Sprintf("%s:%d", s.c.Host, s.c.Port))
	if err != nil {
		return errors.E(op, kind, err, logrus.ErrorLevel)
	}
	return nil
}

func (s server) attachMW() {
	for _, f := range s.mw {
		s.e.Use(f)
	}
}

func (s server) setupRoutes() {
	s.e.Add("POST", "/user/add", api.HandlerAddUser())
}
