package server

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sawalreverr/bebastukar-be/config"
	"github.com/sawalreverr/bebastukar-be/internal/database"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
	gr   *echo.Group
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	app := echo.New()
	app.Validator = &CustomValidator{validator: validator.New()}

	group := app.Group("/api/v1")

	return &echoServer{
		app:  app,
		db:   db,
		conf: conf,
		gr:   group,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	// Healthy Check
	s.app.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Public
	s.publicHttpHandler()

	// Authenticate
	s.authHttpHandler()

	// User
	s.userHttpHandler()

	serverPORT := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverPORT))
}
