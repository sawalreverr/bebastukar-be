package usecase

import (
	"github.com/google/uuid"
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/entity"
	"github.com/sawalreverr/bebastukar-be/internal/repository"
	"github.com/sawalreverr/bebastukar-be/pkg"
)

type discussionUsecase struct {
	discussionRepository repository.DiscussionRepository
}

func NewDiscussionUsecase(discussionRepo repository.DiscussionRepository) DiscussionUsecase {
	return &discussionUsecase{discussionRepository: discussionRepo}
}

func (u *discussionUsecase) CreateDiscussion(discussion dto.DiscussionCredential) (*dto.DiscussionResponse, error) {
	newDiscussion := entity.Discussions{
		ID:      uuid.NewString(),
		UserID:  discussion.AuthorID,
		Content: discussion.Content,
	}

	dataDiscussion, err := u.discussionRepository.CreateDiscussion(newDiscussion)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	for _, url := range discussion.Images {
		newImage := entity.DiscussionImages{
			ID:           uuid.NewString(),
			DiscussionID: dataDiscussion.ID,
			ImageURL:     url,
		}

		_, err := u.discussionRepository.AddImage(newImage)
		if err != nil {
			return nil, pkg.ErrStatusInternalError
		}
	}

	discussionResponse := dto.DiscussionResponse{
		ID:        dataDiscussion.ID,
		AuthorID:  dataDiscussion.UserID,
		Content:   dataDiscussion.Content,
		Images:    discussion.Images,
		CreatedAt: dataDiscussion.CreatedAt,
	}

	return &discussionResponse, nil
}

func (u *discussionUsecase) EditDiscussion(discussionID string, discussion dto.DiscussionCredential) error {
	discussionFound, err := u.discussionRepository.FindDiscussionByID(discussionID)
	if err != nil {
		return pkg.ErrRecordNotFound
	}

	// only edit content
	discussionFound.Content = discussion.Content

	err = u.discussionRepository.UpdateDiscussion(*discussionFound)
	if err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (u *discussionUsecase) DeleteDiscussion(discussionID string) error {
	discussionFound, err := u.discussionRepository.FindDiscussionByID(discussionID)
	if err != nil {
		return pkg.ErrRecordNotFound
	}

	err = u.discussionRepository.DeleteDiscussion(discussionFound.ID, discussionFound.UserID)
	if err != nil {
		return pkg.ErrStatusInternalError
	}

	err = u.discussionRepository.DeleteAllImage(discussionFound.ID)
	if err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (u *discussionUsecase) GetAllDiscussionFromUser(userID string) (*[]dto.DiscussionResponse, error) {
	discussionFound, err := u.discussionRepository.FindDiscussionByUserID(userID)
	if err != nil {
		return nil, pkg.ErrRecordNotFound
	}

	var discussions []dto.DiscussionResponse
	for _, discus := range *discussionFound {
		data := dto.DiscussionResponse{
			ID:        discus.ID,
			AuthorID:  discus.UserID,
			Content:   discus.Content,
			CreatedAt: discus.CreatedAt,
			UpdatedAt: discus.UpdatedAt,
		}

		images, _ := u.discussionRepository.FindAllImage(discus.ID)
		if images != nil {
			data.Images = images
		}

		discussions = append(discussions, data)
	}

	return &discussions, nil
}

func (u *discussionUsecase) GetDiscussionFromID(discussionID string) (*dto.DiscussionResponse, error) {
	discussionFound, err := u.discussionRepository.FindDiscussionByID(discussionID)
	if err != nil {
		return nil, pkg.ErrRecordNotFound
	}

	imageURLs, _ := u.discussionRepository.FindAllImage(discussionFound.ID)
	data := dto.DiscussionResponse{
		ID:        discussionFound.ID,
		AuthorID:  discussionFound.UserID,
		Content:   discussionFound.Content,
		Images:    imageURLs,
		CreatedAt: discussionFound.CreatedAt,
		UpdatedAt: discussionFound.UpdatedAt,
	}

	return &data, nil
}

func (u *discussionUsecase) GetAllDiscussion(page int, limit int, sortBy string, sortType string) (*dto.DiscussionPaginationResponse, error) {
	var discussionResponses []dto.DiscussionResponse
	discussions, err := u.discussionRepository.FindAllDiscussion(page, limit, sortBy, sortType)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	totalCount, err := u.discussionRepository.CountAllDiscussions()
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	for _, discuss := range *discussions {
		imageURLs, _ := u.discussionRepository.FindAllImage(discuss.ID)

		discussResp := dto.DiscussionResponse{
			ID:        discuss.ID,
			AuthorID:  discuss.UserID,
			Content:   discuss.Content,
			Images:    imageURLs,
			CreatedAt: discuss.CreatedAt,
			UpdatedAt: discuss.UpdatedAt,
		}
		discussionResponses = append(discussionResponses, discussResp)
	}

	paginationResponse := dto.DiscussionPaginationResponse{
		TotalCount:  totalCount,
		Page:        page,
		Limit:       limit,
		Discussions: discussionResponses,
	}

	return &paginationResponse, nil
}

func (u *discussionUsecase) CreateCommentDiscussion(comment dto.DiscussionCommentCredential) (*dto.DiscussionCommentResponse, error) {
	discussionFound, err := u.discussionRepository.FindDiscussionByID(comment.DiscussionID)
	if err != nil {
		return nil, pkg.ErrDiscussionNotFound
	}

	newComment := entity.DiscussionComments{
		ID:           uuid.NewString(),
		DiscussionID: discussionFound.ID,
		UserID:       comment.AuthorID,
		Comment:      comment.Comment,
	}

	dataComment, err := u.discussionRepository.AddComment(newComment)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	commentResponse := dto.DiscussionCommentResponse{
		ID:           dataComment.ID,
		AuthorID:     dataComment.UserID,
		DiscussionID: dataComment.DiscussionID,
		Comment:      dataComment.Comment,
		CreatedAt:    dataComment.CreatedAt,
	}

	return &commentResponse, nil
}

func (u *discussionUsecase) EditCommentDiscussion(discussionCommentID string, comment dto.DiscussionCommentCredential) error {
	_, err := u.discussionRepository.FindDiscussionByID(comment.DiscussionID)
	if err != nil {
		return pkg.ErrDiscussionNotFound
	}

	commentFound, err := u.discussionRepository.FindCommentByID(discussionCommentID)
	if err != nil {
		return pkg.ErrCommentNotFound
	}

	commentFound.Comment = comment.Comment

	if commentFound.UserID != comment.AuthorID {
		return pkg.ErrNoPrivilege
	}

	err = u.discussionRepository.UpdateComment(*commentFound)
	if err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (u *discussionUsecase) DeleteCommentDiscussion(discussionID string, discussionCommentID string, userID string) error {
	discussionFound, err := u.discussionRepository.FindDiscussionByID(discussionID)
	if err != nil {
		return pkg.ErrDiscussionNotFound
	}

	commentFound, err := u.discussionRepository.FindCommentByID(discussionCommentID)
	if err != nil {
		return pkg.ErrCommentNotFound
	}

	if commentFound.UserID != userID {
		return pkg.ErrNoPrivilege
	}

	err = u.discussionRepository.DeleteComment(commentFound.ID, discussionFound.ID, userID)
	if err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (u *discussionUsecase) CreateReplyCommentDiscussion(replyComment dto.DiscussionReplyCommentCredential) (*dto.DiscussionReplyCommentResponse, error) {
	discussionFound, err := u.discussionRepository.FindDiscussionByID(replyComment.DiscussionID)
	if err != nil {
		return nil, pkg.ErrDiscussionNotFound
	}

	commentFound, err := u.discussionRepository.FindCommentByID(replyComment.DiscussionCommentID)
	if err != nil {
		return nil, pkg.ErrCommentNotFound
	}

	newReplyComment := entity.DiscussionReplyComments{
		ID:                  uuid.NewString(),
		DiscussionCommentID: commentFound.ID,
		DiscussionID:        discussionFound.ID,
		UserID:              replyComment.AuthorID,
		Comment:             replyComment.ReplyComment,
	}

	dataReplyComment, err := u.discussionRepository.AddReplyComment(newReplyComment)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	replyCommentResponse := dto.DiscussionReplyCommentResponse{
		ID:                  dataReplyComment.ID,
		AuthorID:            dataReplyComment.UserID,
		DiscussionID:        dataReplyComment.DiscussionID,
		DiscussionCommentID: dataReplyComment.DiscussionCommentID,
		ReplyComment:        dataReplyComment.Comment,
		CreatedAt:           dataReplyComment.CreatedAt,
	}

	return &replyCommentResponse, nil
}

func (u *discussionUsecase) EditReplyCommentDiscussion(discussionReplyCommentID string, replyComment dto.DiscussionReplyCommentCredential) error {
	_, err := u.discussionRepository.FindDiscussionByID(replyComment.DiscussionID)
	if err != nil {
		return pkg.ErrDiscussionNotFound
	}

	_, err = u.discussionRepository.FindCommentByID(replyComment.DiscussionCommentID)
	if err != nil {
		return pkg.ErrCommentNotFound
	}

	replyCommentFound, err := u.discussionRepository.FindReplyCommentByID(discussionReplyCommentID)
	if err != nil {
		return pkg.ErrReplyCommentNotFound
	}

	if replyCommentFound.UserID != replyComment.AuthorID {
		return pkg.ErrNoPrivilege
	}

	replyCommentFound.Comment = replyComment.ReplyComment

	err = u.discussionRepository.UpdateReplyComment(*replyCommentFound)
	if err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (u *discussionUsecase) DeleteReplyCommentDiscussion(discussionID string, discussionCommentID string, discussionReplyCommentID string, userID string) error {
	_, err := u.discussionRepository.FindDiscussionByID(discussionID)
	if err != nil {
		return pkg.ErrDiscussionNotFound
	}

	_, err = u.discussionRepository.FindCommentByID(discussionCommentID)
	if err != nil {
		return pkg.ErrCommentNotFound
	}

	replyCommentFound, err := u.discussionRepository.FindReplyCommentByID(discussionReplyCommentID)
	if err != nil {
		return pkg.ErrReplyCommentNotFound
	}

	if replyCommentFound.UserID != userID {
		return pkg.ErrNoPrivilege
	}

	err = u.discussionRepository.DeleteReplyComment(replyCommentFound.ID, replyCommentFound.DiscussionCommentID, replyCommentFound.DiscussionID, replyCommentFound.UserID)
	if err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}
