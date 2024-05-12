package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/entity"
	"github.com/sawalreverr/bebastukar-be/internal/usecase"
	"github.com/sawalreverr/bebastukar-be/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateDiscussion(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)

	authorID := uuid.NewString()
	content := "Test content"
	images := []string{"image1.jpg", "image2.jpg"}

	mockRepo.On("CreateDiscussion", mock.AnythingOfType("entity.Discussions")).Return(&entity.Discussions{ID: "1", UserID: authorID, Content: content, CreatedAt: time.Now()}, nil)
	mockRepo.On("AddImage", mock.AnythingOfType("entity.DiscussionImages")).Return(nil, nil)

	req := dto.DiscussionCredential{
		AuthorID: authorID,
		Content:  content,
		Images:   images,
	}

	resp, err := uc.CreateDiscussion(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "1", resp.ID)
	assert.Equal(t, authorID, resp.AuthorID)
	assert.Equal(t, content, resp.Content)
	assert.Equal(t, images, resp.Images)
}

func TestEditDiscussion(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)

	discussionID := uuid.NewString()
	content := "Updated content"

	mockRepo.On("FindDiscussionByID", discussionID).Return(&entity.Discussions{ID: discussionID, UserID: "userID", Content: "Old content", CreatedAt: time.Now()}, nil)
	mockRepo.On("UpdateDiscussion", mock.AnythingOfType("entity.Discussions")).Return(nil)

	req := dto.DiscussionCredential{
		Content: content,
	}

	err := uc.EditDiscussion(discussionID, req)

	assert.NoError(t, err)
}

func TestDeleteDiscussion(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)

	discussionID := uuid.NewString()

	mockRepo.On("FindDiscussionByID", discussionID).Return(&entity.Discussions{ID: discussionID, UserID: "userID"}, nil)
	mockRepo.On("DeleteDiscussion", discussionID, "userID").Return(nil)
	mockRepo.On("DeleteAllImage", discussionID).Return(nil)

	err := uc.DeleteDiscussion(discussionID)

	assert.NoError(t, err)
}

func TestGetAllDiscussionFromUser(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)

	userID := uuid.NewString()

	discussions := []entity.Discussions{
		{ID: "1", UserID: userID, Content: "Discussion 1", CreatedAt: time.Now()},
		{ID: "2", UserID: userID, Content: "Discussion 2", CreatedAt: time.Now()},
	}

	mockRepo.On("FindDiscussionByUserID", userID).Return(&discussions, nil)
	mockRepo.On("FindAllImage", mock.AnythingOfType("string")).Return([]string{"image1.jpg", "image2.jpg"}, nil)
	mockRepo.On("FindAllComment", mock.AnythingOfType("string")).Return(&[]entity.DiscussionComments{}, nil)

	resp, err := uc.GetAllDiscussionFromUser(userID)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, *resp, len(discussions))
}

func TestGetDiscussionFromID(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)

	discussionID := uuid.NewString()

	discussion := entity.Discussions{
		ID:        discussionID,
		UserID:    uuid.NewString(),
		Content:   "Test discussion",
		CreatedAt: time.Now(),
	}

	mockRepo.On("FindDiscussionByID", discussionID).Return(&discussion, nil)
	mockRepo.On("FindAllImage", discussionID).Return([]string{"image1.jpg", "image2.jpg"}, nil)
	mockRepo.On("FindAllComment", mock.AnythingOfType("string")).Return(&[]entity.DiscussionComments{}, nil)

	resp, err := uc.GetDiscussionFromID(discussionID)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, discussionID, resp.ID)
	assert.Equal(t, discussion.Content, resp.Content)
	assert.Equal(t, 2, len(resp.Images))
}

func TestGetAllDiscussion(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)

	page := 1
	limit := 10
	sortBy := "createdAt"
	sortType := "desc"

	discussions := []entity.Discussions{
		{ID: "1", UserID: uuid.NewString(), Content: "Discussion 1", CreatedAt: time.Now()},
		{ID: "2", UserID: uuid.NewString(), Content: "Discussion 2", CreatedAt: time.Now()},
	}

	mockRepo.On("FindAllDiscussion", page, limit, sortBy, sortType).Return(&discussions, nil)
	mockRepo.On("CountAllDiscussions").Return(len(discussions), nil)
	mockRepo.On("FindAllImage", mock.AnythingOfType("string")).Return([]string{"image1.jpg", "image2.jpg"}, nil)
	mockRepo.On("FindAllComment", mock.AnythingOfType("string")).Return(&[]entity.DiscussionComments{}, nil)

	resp, err := uc.GetAllDiscussion(page, limit, sortBy, sortType)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Discussions, len(discussions))
	assert.Equal(t, len(discussions), resp.TotalCount)
}

func TestCreateCommentDiscussion(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)

	authorID := uuid.NewString()
	discussionID := uuid.NewString()
	comment := "Test comment"

	mockRepo.On("FindDiscussionByID", discussionID).Return(&entity.Discussions{ID: discussionID, UserID: "userID"}, nil)
	mockRepo.On("AddComment", mock.AnythingOfType("entity.DiscussionComments")).Return(&entity.DiscussionComments{ID: "1", UserID: authorID, DiscussionID: discussionID, Comment: comment, CreatedAt: time.Now()}, nil)

	req := dto.DiscussionCommentCredential{
		AuthorID:     authorID,
		DiscussionID: discussionID,
		Comment:      comment,
	}

	resp, err := uc.CreateCommentDiscussion(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, authorID, resp.AuthorID)
	assert.Equal(t, discussionID, resp.DiscussionID)
	assert.Equal(t, comment, resp.Comment)
}

func TestEditCommentDiscussion(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)
	replyComment := "Updated comment"
	commentID := uuid.NewString()
	discussionID := uuid.NewString()

	mockRepo.On("FindCommentByID", commentID).Return(&entity.DiscussionComments{ID: commentID, UserID: "userID", Comment: "Old comment", CreatedAt: time.Now()}, nil)
	mockRepo.On("FindDiscussionByID", discussionID).Return(&entity.Discussions{ID: discussionID}, nil)
	mockRepo.On("UpdateComment", mock.AnythingOfType("entity.DiscussionComments")).Return(nil)

	req := dto.DiscussionCommentCredential{
		AuthorID:     "userID",
		DiscussionID: discussionID,
		Comment:      replyComment,
	}

	err := uc.EditCommentDiscussion(commentID, req)

	assert.NoError(t, err)
}

func TestDeleteCommentDiscussion(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)
	discussionID := uuid.NewString()
	commentID := uuid.NewString()
	userID := uuid.NewString()

	mockRepo.On("FindDiscussionByID", discussionID).Return(&entity.Discussions{ID: discussionID}, nil)
	mockRepo.On("FindCommentByID", commentID).Return(&entity.DiscussionComments{ID: commentID, UserID: userID}, nil)
	mockRepo.On("DeleteComment", commentID, discussionID, userID).Return(nil)

	err := uc.DeleteCommentDiscussion(discussionID, commentID, userID)

	assert.NoError(t, err)
}

func TestGetAllReplyCommentFromComment(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)
	commentID := uuid.NewString()

	replyComments := []entity.DiscussionReplyComments{
		{ID: "1", UserID: "userID", DiscussionID: uuid.NewString(), DiscussionCommentID: commentID, Comment: "Reply 1", CreatedAt: time.Now()},
		{ID: "2", UserID: "userID", DiscussionID: uuid.NewString(), DiscussionCommentID: commentID, Comment: "Reply 2", CreatedAt: time.Now()},
	}

	mockRepo.On("FindAllReplyComment", commentID).Return(&replyComments, nil)

	resp, err := uc.GetAllReplyCommentFromComment(commentID)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, *resp, len(replyComments))
}

func TestCreateReplyCommentDiscussion(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)
	authorID := uuid.NewString()
	discussionID := uuid.NewString()
	commentID := uuid.NewString()
	replyComment := "Test reply comment"

	mockRepo.On("FindDiscussionByID", discussionID).Return(&entity.Discussions{ID: discussionID, UserID: "userID"}, nil)
	mockRepo.On("FindCommentByID", commentID).Return(&entity.DiscussionComments{ID: commentID, UserID: "userID"}, nil)
	mockRepo.On("AddReplyComment", mock.AnythingOfType("entity.DiscussionReplyComments")).Return(&entity.DiscussionReplyComments{ID: "1", UserID: authorID, DiscussionID: discussionID, DiscussionCommentID: commentID, Comment: replyComment, CreatedAt: time.Now()}, nil)

	req := dto.DiscussionReplyCommentCredential{
		AuthorID:            authorID,
		DiscussionID:        discussionID,
		DiscussionCommentID: commentID,
		ReplyComment:        replyComment,
	}

	resp, err := uc.CreateReplyCommentDiscussion(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, authorID, resp.AuthorID)
	assert.Equal(t, discussionID, resp.DiscussionID)
	assert.Equal(t, commentID, resp.DiscussionCommentID)
	assert.Equal(t, replyComment, resp.ReplyComment)
}

func TestEditReplyCommentDiscussion(t *testing.T) {
	mockRepo := new(mocks.MockDiscussionRepository)
	uc := usecase.NewDiscussionUsecase(mockRepo)
	replyCommentID := uuid.NewString()
	replyComment := "Updated reply comment"
	discussionID := uuid.NewString()
	commentID := uuid.NewString()

	mockRepo.On("FindDiscussionByID", discussionID).Return(&entity.Discussions{ID: discussionID, UserID: "userID"}, nil)
	mockRepo.On("FindCommentByID", commentID).Return(&entity.DiscussionComments{ID: commentID, UserID: "userID"}, nil)
	mockRepo.On("FindReplyCommentByID", replyCommentID).Return(&entity.DiscussionReplyComments{ID: replyCommentID, UserID: "userID", Comment: "Old reply comment", CreatedAt: time.Now()}, nil)
	mockRepo.On("UpdateReplyComment", mock.AnythingOfType("entity.DiscussionReplyComments")).Return(nil)

	req := dto.DiscussionReplyCommentCredential{
		AuthorID:            "userID",
		DiscussionID:        discussionID,
		DiscussionCommentID: commentID,
		ReplyComment:        replyComment,
	}

	err := uc.EditReplyCommentDiscussion(replyCommentID, req)

	assert.NoError(t, err)
}
