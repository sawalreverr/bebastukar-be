package usecase

import (
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/entity"
)

type UserUsecase interface {
	RegisterUser(user dto.UserCredential) (*entity.User, error)
	LoginUser(email string, password string) (string, error)
	UpdateUser(userID string, user dto.UpdateUser) error
	UpdateUserAvatar(userID, imageUrl string) error
	FindUserByID(userID string) (*entity.User, error)
	FindAllUser(page int, limit int, sortBy string, sortType string) (*[]entity.User, error)
	DeleteUser(userID string) error
}

type DiscussionUsecase interface {
	// Discussion
	CreateDiscussion(discussion dto.DiscussionCredential) (*dto.DiscussionResponse, error)
	EditDiscussion(discussionID string, discussion dto.DiscussionCredential) error
	DeleteDiscussion(discussionID string) error

	// Discussion Find
	GetAllDiscussionFromUser(userID string) (*[]dto.DiscussionResponse, error)
	GetDiscussionFromID(discussionID string) (*dto.DiscussionResponse, error)
	GetAllDiscussion(page int, limit int, sortBy string, sortType string) (*dto.DiscussionPaginationResponse, error)

	// Discussion Comment
	CreateCommentDiscussion(comment dto.DiscussionCommentCredential) (*dto.DiscussionCommentResponse, error)
	EditCommentDiscussion(discussionCommentID string, comment dto.DiscussionCommentCredential) error
	DeleteCommentDiscussion(discussionID string, discussionCommentID string, userID string) error

	// Discussion Comment Find
	GetAllCommentFromDiscussion(discussionID string) (*[]dto.DiscussionCommentResponse, error)
	GetAllCommentFromDiscussionPublic(discussionID string, discussionCommentID string) (*[]dto.DiscussionCommentResponse, error)

	// Discussion Reply Comment
	CreateReplyCommentDiscussion(replyComment dto.DiscussionReplyCommentCredential) (*dto.DiscussionReplyCommentResponse, error)
	EditReplyCommentDiscussion(discussionReplyCommentID string, replyComment dto.DiscussionReplyCommentCredential) error
	DeleteReplyCommentDiscussion(discussionID string, discussionCommentID string, discussionReplyCommentID string, userID string) error

	// Discussion Reply Comment Find
	GetAllReplyCommentFromComment(discussionCommentID string) (*[]dto.DiscussionReplyCommentResponse, error)

	// ChatBot with Google AI
	// TODO
	// Create history chat from userID to database
	// Find all history chat from userID
}
