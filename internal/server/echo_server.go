package server

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sawalreverr/bebastukar-be/config"
	"github.com/sawalreverr/bebastukar-be/internal/database"
	"github.com/sawalreverr/bebastukar-be/internal/handler"
	"github.com/sawalreverr/bebastukar-be/internal/repository"
	"github.com/sawalreverr/bebastukar-be/internal/usecase"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
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

	return &echoServer{
		app:  app,
		db:   db,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	// Healthy Check
	s.app.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Authenticate
	s.authHttpHandler()

	serverPORT := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverPORT))
}

func (s *echoServer) authHttpHandler() {
	// Depedency
	userRepository := repository.NewUserRepository(s.db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	authHandler := handler.NewAuthHandler(userUsecase)

	// Route
	auth := s.app.Group("/api/v1")
	auth.POST("/register", authHandler.RegisterHandler)
	auth.POST("/login", authHandler.LoginHandler)
}
