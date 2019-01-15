package domain

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// News represent entity of the news
type News struct {
	gorm.Model
	Title string `json:"title"`
	Slug  string `json:"slug"`
}
