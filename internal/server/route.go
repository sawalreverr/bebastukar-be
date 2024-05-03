package server

import (
	"github.com/sawalreverr/bebastukar-be/internal/handler"
	"github.com/sawalreverr/bebastukar-be/internal/middleware"
	"github.com/sawalreverr/bebastukar-be/internal/repository"
	"github.com/sawalreverr/bebastukar-be/internal/usecase"
)

func (s *echoServer) publicHttpHandler() {
	// Depedency
	userRepository := repository.NewUserRepository(s.db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	// Find user by id
	s.gr.GET("/users/:id", userHandler.FindUser)
}

func (s *echoServer) authHttpHandler() {
	// Depedency
	userRepository := repository.NewUserRepository(s.db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	authHandler := handler.NewAuthHandler(userUsecase)

	// Route
	s.gr.POST("/register", authHandler.RegisterHandler)
	s.gr.POST("/login", authHandler.LoginHandler)
}

func (s *echoServer) userHttpHandler() {
	// Depedency
	userRepository := repository.NewUserRepository(s.db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	// Router
	user := s.gr.Group("", middleware.JWTMiddleware)
	user.GET("/users/profile", userHandler.ProfileGet)
	user.POST("/users/profile", userHandler.ProfileUpdate)
	user.POST("/users/uploadAvatar", userHandler.UploadAvatar)
}
