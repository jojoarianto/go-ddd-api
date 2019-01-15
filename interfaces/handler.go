package interfaces

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jojoarianto/go-ddd-api/application"
	"github.com/julienschmidt/httprouter"
)

// Run start server
func Run(port int) error {
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	r := httprouter.New()
	r.GET("/api/v1/news/:news_id", getNews)

	return r
}

// =============================
//    NEWS
// =============================

// Handler for get all news
func getNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newsID, err := strconv.Atoi(ps.ByName("news_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, "invalid parameter")
		return
	}

	news, err := application.GetNews(newsID)
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get news")
		return
	}

	JSON(w, http.StatusOK, news)
}

// =============================
//    TOPIC
// =============================
