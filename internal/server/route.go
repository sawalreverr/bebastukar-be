package server

import (
	"github.com/sawalreverr/bebastukar-be/internal/chatbot"
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
	s.gr.GET("/discussion/user/:userID", discussionHandler.FindAllDiscussionUserHandler)

	// Find all discussion pagination
	s.gr.GET("/discussions", discussionHandler.FindAllDiscussion)

	// Find all discussion comment
	s.gr.GET("/discussion/:id/:commentID", discussionHandler.FindAllDiscussionCommentHandler)
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
	user.GET("/users", userHandler.ProfileGet)
	user.PUT("/users", userHandler.ProfileUpdate)
	user.POST("/users/uploadAvatar", userHandler.UploadAvatar)
}

func (s *echoServer) discussionHttpHandler() {
	// Dependecy
	discussionRepository := repository.NewDiscussionRepository(s.db)
	discussionUsecase := usecase.NewDiscussionUsecase(discussionRepository)
	discussionHandler := handler.NewDiscussionHandler(discussionUsecase)

	// Discussion Router
	discussion := s.gr.Group("", middleware.JWTMiddleware)
	discussion.GET("/discussion", discussionHandler.GetAllDiscussionFromProfile)
	discussion.POST("/discussion", discussionHandler.NewDiscussionHandler)
	discussion.PUT("/discussion/:id", discussionHandler.EditDiscussionhandler)
	discussion.DELETE("/discussion/:id", discussionHandler.DeleteDiscussionhandler)

	// Discussion Comment Router
	discussion.POST("/discussion/:id/comment", discussionHandler.AddDiscussionCommentHandler)
	discussion.PUT("/discussion/:id/:commentID", discussionHandler.EditDiscussionCommentHandler)
	discussion.DELETE("/discussion/:id/:commentID", discussionHandler.DeleteDiscussionCommentHandler)

	// Discussion Reply Comment Router
	discussion.POST("/discussion/:id/:commentID/reply", discussionHandler.AddDiscussionReplyCommentHandler)
	discussion.PUT("/discussion/:id/:commentID/:replyCommentID", discussionHandler.EditDiscussionReplyCommentHandler)
	discussion.DELETE("/discussion/:id/:commentID/:replyCommentID", discussionHandler.DeleteDiscussionReplyCommentHandler)
}

func (s *echoServer) chatbotHttpHandler() {
	// This is only for our first version of my project will be updated later
	chatBotHandler := chatbot.NewDiscussionHandler()

	chatBotGroup := s.gr.Group("", middleware.JWTMiddleware)
	chatBotGroup.POST("/chatbot", chatBotHandler.QuestionHandler)
}
