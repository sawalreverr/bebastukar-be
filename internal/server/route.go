package server

import (
	"github.com/sawalreverr/bebastukar-be/internal/handler"
	"github.com/sawalreverr/bebastukar-be/internal/middleware"
	"github.com/sawalreverr/bebastukar-be/internal/repository"
	"github.com/sawalreverr/bebastukar-be/internal/usecase"
)

func (s *echoServer) publicHttpHandler() {
	// Depedency user
	userRepository := repository.NewUserRepository(s.db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	// Dependency discussion
	discussionRepository := repository.NewDiscussionRepository(s.db)
	discussionUsecase := usecase.NewDiscussionUsecase(discussionRepository)
	discussionHandler := handler.NewDiscussionHandler(discussionUsecase)

	// Find user by id
	s.gr.GET("/users/:id", userHandler.FindUser)

	// Find discussion by id
	s.gr.GET("/discussion/:id", discussionHandler.FindDiscussionByID)

	// Find all discussion from user
	s.gr.GET("/discussion/user/:userid", discussionHandler.FindAllDiscussionUserHandler)
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
	user.PUT("/users/profile", userHandler.ProfileUpdate)
	user.POST("/users/uploadAvatar", userHandler.UploadAvatar)
}

func (s *echoServer) discussionHttpHandler() {
	// Dependecy
	discussionRepository := repository.NewDiscussionRepository(s.db)
	discussionUsecase := usecase.NewDiscussionUsecase(discussionRepository)
	discussionHandler := handler.NewDiscussionHandler(discussionUsecase)

	// Router
	discussion := s.gr.Group("", middleware.JWTMiddleware)
	discussion.GET("/discussion", discussionHandler.GetAllDiscussionFromProfile)
	discussion.POST("/discussion", discussionHandler.NewDiscussionHandler)
	discussion.PUT("/discussion/:id", discussionHandler.EditDiscussionhandler)
	discussion.DELETE("/discussion/:id", discussionHandler.DeleteDiscussionhandler)
}
