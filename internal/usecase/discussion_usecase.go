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

	// if imageURLs != nil {
	// 	data.Images = *imageURLs
	// }

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
