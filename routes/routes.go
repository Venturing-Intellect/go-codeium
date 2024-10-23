// routes/routes.go
package routes

import (
	"customer-feedback-api/controllers"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, controller *controllers.FeedbackController) {
	mux.HandleFunc("/feedback", controller.GetAllFeedback)
	mux.HandleFunc("/feedback/create", controller.CreateFeedback)
}
