package persistence

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jojoarianto/go-ddd-api/domain"
	"github.com/jojoarianto/go-ddd-api/domain/repository"
)

// NewsRepositoryImpl Implements repository.NewsRepository
type NewsRepositoryImpl struct {
	Conn *gorm.DB
}

// NewNewsRepositoryWithRDB returns initialized NewsRepositoryImpl
func NewNewsRepositoryWithRDB(conn *gorm.DB) repository.NewsRepository {
	return &NewsRepositoryImpl{Conn: conn}
}

// Get news by id return domain.news
func (r *NewsRepositoryImpl) Get(id int) (*domain.News, error) {
	news := &domain.News{}
	if err := r.Conn.First(&news, id).Error; err != nil {
		return nil, err
	}
	return news, nil
}
