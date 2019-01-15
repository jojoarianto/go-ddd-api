package interfaces

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jojoarianto/go-ddd-api/application"
	"github.com/jojoarianto/go-ddd-api/config"
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

	// News Route
	r.GET("/api/v1/news/:news_id", getNews)
	// Topic Route
	r.GET("/api/v1/topic/:topic_id", getTopic)
	// Migration Route
	r.GET("/api/v1/migrate", migrate)

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
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, news)
}

// =============================
//    TOPIC
// =============================

// Handler for get Topic by id
func getTopic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	topicID, err := strconv.Atoi(ps.ByName("topic_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, "invalid parameter")
		return
	}

	topic, err := application.GetTopic(topicID)
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get topic")
		return
	}

	JSON(w, http.StatusOK, topic)
}

// =============================
//    MIGRATE
// =============================

func migrate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := config.DBMigrate()
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to migrate")
		return
	}

	msg := "Success Migrate"
	JSON(w, http.StatusOK, msg)
}
