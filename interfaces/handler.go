package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/jojoarianto/go-ddd-api/application"
	"github.com/jojoarianto/go-ddd-api/config"
	"github.com/jojoarianto/go-ddd-api/domain"
	"github.com/julienschmidt/httprouter"
)

// IsLetter function to check string is aplhanumeric only
var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

// Run start server
func Run(port int) error {
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	r := httprouter.New()

	// Index Route
	r.GET("/", index)
	r.GET("/api/v1", index)

	// News Route
	r.GET("/api/v1/news", getAllNews)
	r.GET("/api/v1/news/:param", getNews)
	r.POST("/api/v1/news", createNews)
	r.DELETE("/api/v1/news/:news_id", removeNews)
	r.PUT("/api/v1/news/:news_id", updateNews)

	// Topic Route
	r.GET("/api/v1/topic", getAllTopic)
	r.GET("/api/v1/topic/:topic_id", getTopic)
	r.POST("/api/v1/topic", createTopic)
	r.DELETE("/api/v1/topic/:topic_id", removeTopic)
	r.PUT("/api/v1/topic/:topic_id", updateTopic)

	// Migration Route
	r.GET("/api/v1/migrate", migrate)

	return r
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	JSON(w, http.StatusOK, "GO DDD API")
}

// =============================
//    NEWS
// =============================

func getNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("param")

	// if param is numeric than search by news_id, otherwise
	// if alphabetic then search by topic.Slug
	newsID, err := strconv.Atoi(param)
	if err != nil {
		// param is alphabetic
		news, err2 := application.GetNewsByTopic(param)
		if err2 != nil {
			Error(w, http.StatusNotFound, err2, err2.Error())
			return
		}

		JSON(w, http.StatusOK, news)
		return
	}

	// param is numeric
	news, err := application.GetNews(newsID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, news)
}

func getAllNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
	status := queryValues.Get("status")

	// if status parameter exist draft|deleted|publish
	if status == "draft" || status == "deleted" || status == "publish" {
		news, err := application.GetAllNewsByFilter(status)
		if err != nil {
			Error(w, http.StatusNotFound, err, err.Error())
			return
		}

		JSON(w, http.StatusOK, news)
		return
	}

	limit := queryValues.Get("limit")
	page := queryValues.Get("page")
	
	// if custom pagination exist news?limit=15&page=2
	if limit != "" && page != "" {
		limit, _ := strconv.Atoi(limit)
		page, _ := strconv.Atoi(page)

		if limit != 0 && page != 0 {
			news, err := application.GetAllNews(limit, page)
			if err != nil {
				Error(w, http.StatusNotFound, err, err.Error())
				return
			}
	
			JSON(w, http.StatusOK, news)
			return
		}
	}

	news, err := application.GetAllNews(15, 1) // 15, 1 default pagination
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, news)
}

func createNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p domain.News
	if err := decoder.Decode(&p); err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	err := application.AddNews(p)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusCreated, nil)
}

func removeNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newsID, err := strconv.Atoi(ps.ByName("news_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.RemoveNews(newsID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, nil)
}

func updateNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p domain.News
	err := decoder.Decode(&p)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	newsID, err := strconv.Atoi(ps.ByName("news_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.UpdateNews(p, newsID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, nil)
}

// =============================
//    TOPIC
// =============================

func getTopic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	topicID, err := strconv.Atoi(ps.ByName("topic_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	topic, err := application.GetTopic(topicID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, topic)
}

func getAllTopic(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	topics, err := application.GetAllTopic()
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, topics)
}

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
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusCreated, nil)
}

func removeTopic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	topicID, err := strconv.Atoi(ps.ByName("topic_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.RemoveTopic(topicID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, nil)
}

func updateTopic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p domain.Topic
	err := decoder.Decode(&p)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	topicID, err := strconv.Atoi(ps.ByName("topic_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.UpdateTopic(p, topicID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
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
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	msg := "Success Migrate"
	JSON(w, http.StatusOK, msg)
}
