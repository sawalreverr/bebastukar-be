package repository

import "github.com/sawalreverr/bebastukar-be/internal/entity"

type UserRepository interface {
	Create(user entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByID(userID string) (*entity.User, error)
	FindAll(page int, limit int, sortBy string, sortType string) (*[]entity.User, error)
	Update(user *entity.User) error
	Delete(userID string) error
}

type DiscussionRepository interface {
	// Discussion
	CreateDiscussion(discussion entity.Discussions) (*entity.Discussions, error)
	UpdateDiscussion(discussion entity.Discussions) error
	DeleteDiscussion(discussionID string, userID string) error
	FindDiscussionByID(discussionID string) (*entity.Discussions, error)
	FindDiscussionByUserID(userID string) (*[]entity.Discussions, error)
	FindAllDiscussion() (*[]entity.Discussions, error)

	// Discussion Image
	AddImage(discussionImage entity.DiscussionImages) (*entity.DiscussionImages, error)
	DeleteImage(discussionImageID string, discussionID string) error
	FindAllImage(discussionID string) (*[]string, error)

	// Discussion Comment
	AddComment(comment entity.DiscussionComments) (*entity.DiscussionComments, error)
	UpdateComment(comment entity.DiscussionComments) error
	DeleteComment(discussionCommentID string, discussionID string, userID string) error
	FindAllComment(discussionID string) (*[]entity.DiscussionComments, error)

	// Discussion Reply Comment
	AddReplyComment(replyComment entity.DiscussionReplyComments) (*entity.DiscussionReplyComments, error)
	UpdateReplyComment(replyComment entity.DiscussionReplyComments) error
	DeleteReplyComment(discussionReplyCommentID string, discussionCommentID string, userID string) error
	FindAllReplyComment(discussionCommentID string) (*[]entity.DiscussionReplyComments, error)
}
