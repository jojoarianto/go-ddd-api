package interfaces

import (
	"encoding/json"
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
	r.GET("/api/v1/news", getAllNews)
	r.GET("/api/v1/news/:news_id", getNews)
	r.POST("/api/v1/news", createNews)
	// Topic Route
	r.GET("/api/v1/topic", getAllTopic)
	r.GET("/api/v1/topic/:topic_id", getTopic)
	r.POST("/api/v1/topic", createTopic)
	r.DELETE("/api/v1/topic/:topic_id", removeTopic)

	// Migration Route
	r.GET("/api/v1/migrate", migrate)

	return r
}

// =============================
//    NEWS
// =============================

// getNews handler for handler get news by id
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

// getAllNews handler for handler get all news
func getAllNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	news, err := application.GetAllNews()
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get news")
		return
	}

	JSON(w, http.StatusOK, news)
}

// createNews handler for handler create news
func createNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type payload struct {
		Title string `json:"title"`
		Slug  string `json:"slug"`
	}
	var p payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.AddNews(p.Title, p.Slug)
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to create news")
		return
	}

	JSON(w, http.StatusCreated, nil)
}

// =============================
//    TOPIC
// =============================

// getTopic handler for handler get topic by id
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

// getAllTopic handler for handler get all topics
func getAllTopic(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	topics, err := application.GetAllTopic()
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get topics")
		return
	}

	JSON(w, http.StatusOK, topics)
}

// createNews handler for handler remove topic
func createTopic(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	type payload struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	}
	var p payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.AddTopic(p.Name, p.Slug)
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to create topic")
		return
	}

	JSON(w, http.StatusCreated, nil)
}

// removeTopic handler for handler remove topic
func removeTopic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	topicID, err := strconv.Atoi(ps.ByName("topic_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, "invalid parameter")
		return
	}

	err = application.RemoveTopic(topicID)
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to delete topic")
		return
	}

	JSON(w, http.StatusOK, nil)
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
