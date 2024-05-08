package repository

import (
	"fmt"

	"github.com/sawalreverr/bebastukar-be/internal/database"
	"github.com/sawalreverr/bebastukar-be/internal/entity"
)

type discussionRepository struct {
	DB database.Database
}

func NewDiscussionRepository(db database.Database) DiscussionRepository {
	return &discussionRepository{DB: db}
}

// Discussion Repository
func (r *discussionRepository) CreateDiscussion(discussion entity.Discussions) (*entity.Discussions, error) {
	if err := r.DB.GetDB().Create(&discussion).Error; err != nil {
		return nil, err
	}

	return &discussion, nil
}

func (r *discussionRepository) UpdateDiscussion(discussion entity.Discussions) error {
	if err := r.DB.GetDB().Save(&discussion).Error; err != nil {
		return err
	}

	return nil
}

func (r *discussionRepository) DeleteDiscussion(discussionID string, userID string) error {
	var discussion *entity.Discussions
	if err := r.DB.GetDB().Where("id = ? AND user_id = ?", discussionID, userID).Delete(&discussion).Error; err != nil {
		return err
	}

	return nil
}

func (r *discussionRepository) FindDiscussionByID(discussionID string) (*entity.Discussions, error) {
	var discussion entity.Discussions
	if err := r.DB.GetDB().Where("id = ?", discussionID).First(&discussion).Error; err != nil {
		return nil, err
	}

	return &discussion, nil
}

func (r *discussionRepository) FindDiscussionByUserID(userID string) (*[]entity.Discussions, error) {
	var userDiscussions []entity.Discussions
	if err := r.DB.GetDB().Where("user_id = ?", userID).Find(&userDiscussions).Error; err != nil {
		return nil, err
	}

	return &userDiscussions, nil
}

func (r *discussionRepository) FindAllDiscussion(page int, limit int, sortBy string, sortType string) (*[]entity.Discussions, error) {
	var discussions *[]entity.Discussions

	db := r.DB.GetDB()
	offset := (page - 1) * limit

	if sortBy != "" {
		sort := fmt.Sprintf("%s %s", sortBy, sortType)
		db = db.Order(sort)
	}

	if err := db.Offset(offset).Limit(limit).Find(&discussions).Error; err != nil {
		return nil, err
	}

	return discussions, nil
}

func (r *discussionRepository) CountAllDiscussions() (int, error) {
	var totalCount int64

	if err := r.DB.GetDB().Model(&entity.Discussions{}).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	return int(totalCount), nil
}

// Discussion Image Repository
func (r *discussionRepository) AddImage(discussionImage entity.DiscussionImages) (*entity.DiscussionImages, error) {
	if err := r.DB.GetDB().Create(&discussionImage).Error; err != nil {
		return nil, err
	}

	return &discussionImage, nil
}

func (r *discussionRepository) DeleteImage(discussionImageID string, discussionID string) error {
	var discussionImage entity.DiscussionImages
	if err := r.DB.GetDB().Where("id = ? AND discussion_id = ?", discussionImageID, discussionID).Delete(&discussionImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *discussionRepository) DeleteAllImage(discussionID string) error {
	if err := r.DB.GetDB().Where("discussion_id = ?", discussionID).Delete(&entity.DiscussionImages{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *discussionRepository) FindAllImage(discussionID string) ([]string, error) {
	var discussionImages []entity.DiscussionImages
	var imageURLs []string

	if err := r.DB.GetDB().Where("discussion_id = ?", discussionID).Find(&discussionImages).Error; err != nil {
		return nil, err
	}

	for _, discusImage := range discussionImages {
		imageURLs = append(imageURLs, discusImage.ImageURL)
	}

	if len(imageURLs) == 0 {
		return []string{}, nil
	}

	return imageURLs, nil
}

// Discussion Comment Repository
func (r *discussionRepository) AddComment(comment entity.DiscussionComments) (*entity.DiscussionComments, error) {
	if err := r.DB.GetDB().Create(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *discussionRepository) UpdateComment(comment entity.DiscussionComments) error {
	if err := r.DB.GetDB().Save(&comment).Error; err != nil {
		return err
	}

	return nil
}

func (r *discussionRepository) DeleteComment(discussionCommentID string, discussionID string, userID string) error {
	var discussionComment entity.DiscussionComments
	if err := r.DB.GetDB().Where("id = ? AND discussion_id = ? AND user_id = ?", discussionCommentID, discussionID, userID).Delete(&discussionComment).Error; err != nil {
		return err
	}

	return nil
}

func (r *discussionRepository) FindAllComment(discussionID string) (*[]entity.DiscussionComments, error) {
	var discussionComments []entity.DiscussionComments
	if err := r.DB.GetDB().Where("discussion_id = ?", discussionID).Find(&discussionComments).Error; err != nil {
		return nil, err
	}

	return &discussionComments, nil
}

// Discussion Reply Comment Repository
func (r *discussionRepository) AddReplyComment(replyComment entity.DiscussionReplyComments) (*entity.DiscussionReplyComments, error) {
	if err := r.DB.GetDB().Create(&replyComment).Error; err != nil {
		return nil, err
	}

	return &replyComment, nil
}

func (r *discussionRepository) UpdateReplyComment(replyComment entity.DiscussionReplyComments) error {
	if err := r.DB.GetDB().Save(&replyComment).Error; err != nil {
		return err
	}

	return nil
}

func (r *discussionRepository) DeleteReplyComment(discussionReplyCommentID string, discussionCommentID string, userID string) error {
	var discussionReplyComment entity.DiscussionReplyComments
	if err := r.DB.GetDB().Where("id = ? AND discussion_comment_id = ? AND user_id = ?", discussionReplyCommentID, discussionCommentID, userID).Delete(&discussionReplyComment).Error; err != nil {
		return err
	}

	return nil
}

func (r *discussionRepository) FindAllReplyComment(discussionCommentID string) (*[]entity.DiscussionReplyComments, error) {
	var discussionReplyComments []entity.DiscussionReplyComments
	if err := r.DB.GetDB().Where("discussion_comment_id = ?", discussionCommentID).Find(&discussionReplyComments).Error; err != nil {
		return nil, err
	}

	return &discussionReplyComments, nil
}
