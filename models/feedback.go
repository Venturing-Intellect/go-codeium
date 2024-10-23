// models/feedback.go
package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Feedback string `json:"feedback"`
}
