package interfaces

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jojoarianto/go-ddd-api/application"
	router "github.com/takashabe/go-router"
)

// Run start server
func Run(port int) error {
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *router.Router {
	r := router.NewRouter()
	r.Get("/news/:id", getNews)

	return r
}

// Handler for get all news
func getNews(w http.ResponseWriter, r *http.Request, id int) {
	news, err := application.GetNews(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get news")
		return
	}

	JSON(w, http.StatusOK, news)
}
