package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"customer-feedback-api/models"
	"customer-feedback-api/repositories"
	"customer-feedback-api/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateFeedback(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&models.Feedback{})

	repository := repositories.NewFeedbackRepository(db)
	service := services.NewFeedbackService(repository)

	feedback := models.Feedback{Name: "John Doe", Email: "john@example.com", Feedback: "This is a test feedback"}
	jsonFeedback, err := json.Marshal(feedback)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/feedback/create", bytes.NewBuffer(jsonFeedback))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createdFeedback, err := service.CreateFeedback(feedback.Name, feedback.Email, feedback.Feedback)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdFeedback)
	})

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, w.Code)
	}

	var createdFeedback models.Feedback
	err = json.Unmarshal(w.Body.Bytes(), &createdFeedback)
	if err != nil {
		t.Fatal(err)
	}

	if createdFeedback.Name != feedback.Name {
		t.Errorf("Expected name %s but got %s", feedback.Name, createdFeedback.Name)
	}

	if createdFeedback.Email != feedback.Email {
		t.Errorf("Expected email %s but got %s", feedback.Email, createdFeedback.Email)
	}

	if createdFeedback.Feedback != feedback.Feedback {
		t.Errorf("Expected feedback %s but got %s", feedback.Feedback, createdFeedback.Feedback)
	}
}

func TestValidateEmail(t *testing.T) {
	service := services.NewFeedbackService(nil)

	testCases := []struct {
		email    string
		expected bool
	}{
		{"john@example.com", true},
		{"john.doe@example.com", true},
		{"john.doe@example.co.uk", true},
		{"john.doe@example", false},
		{"john.doe@example.", false},
		{"john.doe@example.com.", false},
		{"john.doe@.example.com", false},
	}

	for _, tc := range testCases {
		actual := service.ValidateEmail(tc.email)
		if actual != tc.expected {
			t.Errorf("Expected %v for email %s, but got %v", tc.expected, tc.email, actual)
		}
	}
}
