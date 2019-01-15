package domain

import _ "github.com/jinzhu/gorm/dialects/mysql"

// News represent entity of the news
type News struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}
