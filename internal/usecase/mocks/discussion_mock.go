package mocks

import (
	"github.com/sawalreverr/bebastukar-be/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockDiscussionRepository struct {
	mock.Mock
}

// DeleteImage implements repository.DiscussionRepository.
func (m *MockDiscussionRepository) DeleteImage(discussionImageID string, discussionID string) error {
	panic("unimplemented")
}

func (m *MockDiscussionRepository) CreateDiscussion(discussion entity.Discussions) (*entity.Discussions, error) {
	args := m.Called(discussion)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.Discussions), args.Error(1)
}

func (m *MockDiscussionRepository) AddImage(image entity.DiscussionImages) (*entity.DiscussionImages, error) {
	args := m.Called(image)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.DiscussionImages), args.Error(1)
}

func (m *MockDiscussionRepository) FindDiscussionByID(discussionID string) (*entity.Discussions, error) {
	args := m.Called(discussionID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.Discussions), args.Error(1)
}

func (m *MockDiscussionRepository) UpdateDiscussion(discussion entity.Discussions) error {
	args := m.Called(discussion)
	return args.Error(0)
}

func (m *MockDiscussionRepository) DeleteDiscussion(discussionID string, userID string) error {
	args := m.Called(discussionID, userID)
	return args.Error(0)
}

func (m *MockDiscussionRepository) DeleteAllImage(discussionID string) error {
	args := m.Called(discussionID)
	return args.Error(0)
}

func (m *MockDiscussionRepository) FindDiscussionByUserID(userID string) (*[]entity.Discussions, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*[]entity.Discussions), args.Error(1)
}

func (m *MockDiscussionRepository) FindAllImage(discussionID string) ([]string, error) {
	args := m.Called(discussionID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]string), args.Error(1)
}

func (m *MockDiscussionRepository) FindAllDiscussion(page int, limit int, sortBy string, sortType string) (*[]entity.Discussions, error) {
	args := m.Called(page, limit, sortBy, sortType)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*[]entity.Discussions), args.Error(1)
}

func (m *MockDiscussionRepository) CountAllDiscussions() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockDiscussionRepository) AddComment(comment entity.DiscussionComments) (*entity.DiscussionComments, error) {
	args := m.Called(comment)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.DiscussionComments), args.Error(1)
}

func (m *MockDiscussionRepository) FindCommentByID(commentID string) (*entity.DiscussionComments, error) {
	args := m.Called(commentID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.DiscussionComments), args.Error(1)
}

func (m *MockDiscussionRepository) UpdateComment(comment entity.DiscussionComments) error {
	args := m.Called(comment)
	return args.Error(0)
}

func (m *MockDiscussionRepository) DeleteComment(commentID string, discussionID string, userID string) error {
	args := m.Called(commentID, discussionID, userID)
	return args.Error(0)
}

func (m *MockDiscussionRepository) FindAllComment(discussionID string) (*[]entity.DiscussionComments, error) {
	args := m.Called(discussionID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*[]entity.DiscussionComments), args.Error(1)
}

func (m *MockDiscussionRepository) AddReplyComment(replyComment entity.DiscussionReplyComments) (*entity.DiscussionReplyComments, error) {
	args := m.Called(replyComment)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.DiscussionReplyComments), args.Error(1)
}

func (m *MockDiscussionRepository) FindReplyCommentByID(replyCommentID string) (*entity.DiscussionReplyComments, error) {
	args := m.Called(replyCommentID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*entity.DiscussionReplyComments), args.Error(1)
}

func (m *MockDiscussionRepository) UpdateReplyComment(replyComment entity.DiscussionReplyComments) error {
	args := m.Called(replyComment)
	return args.Error(0)
}

func (m *MockDiscussionRepository) DeleteReplyComment(replyCommentID string, discussionCommentID string, discussionID string, userID string) error {
	args := m.Called(replyCommentID, discussionCommentID, discussionID, userID)
	return args.Error(0)
}

func (m *MockDiscussionRepository) FindAllReplyComment(discussionCommentID string) (*[]entity.DiscussionReplyComments, error) {
	args := m.Called(discussionCommentID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*[]entity.DiscussionReplyComments), args.Error(1)
}
