// services/feedback_service.go
package services

import (
	"customer-feedback-api/models"
	"customer-feedback-api/repositories"
	"errors"
	"regexp"
)

type FeedbackService struct {
	repository *repositories.FeedbackRepository
}

func NewFeedbackService(repository *repositories.FeedbackRepository) *FeedbackService {
	return &FeedbackService{repository: repository}
}

func (s *FeedbackService) ValidateEmail(email string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)+$`)
	return pattern.MatchString(email)
}

func (s *FeedbackService) CreateFeedback(name, email string, feedback string) (*models.Feedback, error) {
	if !s.ValidateEmail(email) {
		return nil, errors.New("invalid email format")
	}
	return s.repository.CreateFeedback(name, email, feedback)
}

func (s *FeedbackService) GetAllFeedback() ([]models.Feedback, error) {
	return s.repository.GetAllFeedback()
}
