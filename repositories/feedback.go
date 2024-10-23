// repositories/feedback_repository.go
package repositories

import (
	"customer-feedback-api/models"

	"gorm.io/gorm"
)

type FeedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) *FeedbackRepository {
	return &FeedbackRepository{db: db}
}

func (r *FeedbackRepository) CreateFeedback(name, email, feedback string) (*models.Feedback, error) {
	newFeedback := models.Feedback{Name: name, Email: email, Feedback: feedback}
	result := r.db.Create(&newFeedback)
	return &newFeedback, result.Error
}

func (r *FeedbackRepository) GetAllFeedback() ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	result := r.db.Find(&feedbacks)
	return feedbacks, result.Error
}
