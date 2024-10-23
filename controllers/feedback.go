// controllers/feedback_controller.go
package controllers

import (
	"customer-feedback-api/models"
	"customer-feedback-api/services"
	"encoding/json"
	"net/http"
)

type FeedbackController struct {
	service *services.FeedbackService
}

func NewFeedbackController(service *services.FeedbackService) *FeedbackController {
	return &FeedbackController{service: service}
}

func (c *FeedbackController) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback models.Feedback
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdFeedback, err := c.service.CreateFeedback(feedback.Name, feedback.Email, feedback.Feedback)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdFeedback)
}

func (c *FeedbackController) GetAllFeedback(w http.ResponseWriter, r *http.Request) {
	feedbacks, err := c.service.GetAllFeedback()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(feedbacks)
}
