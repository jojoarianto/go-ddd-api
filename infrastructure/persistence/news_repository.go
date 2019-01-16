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
	if err := r.Conn.Preload("Topic").First(&news, id).Error; err != nil {
		return nil, err
	}
	return news, nil
}

// GetAll News return all domain.news
func (r *NewsRepositoryImpl) GetAll() ([]domain.News, error) {
	news := []domain.News{}
	if err := r.Conn.Preload("Topic").Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}

// Save to add news
func (r *NewsRepositoryImpl) Save(news *domain.News) error {
	if err := r.Conn.Save(&news).Error; err != nil {
		return err
	}

	return nil
}

// Remove to delete news by id
func (r *NewsRepositoryImpl) Remove(id int) error {
	tx := r.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	news := domain.News{}
	if err := tx.First(&news, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	news.Status = "deleted"
	if err := tx.Save(&news).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&news).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Update is update news
func (r *NewsRepositoryImpl) Update(news *domain.News) error {
	if err := r.Conn.Model(&news).UpdateColumns(domain.News{Title: news.Title, Slug: news.Slug, Content: news.Content, Status: news.Status, Topic: news.Topic}).Error; err != nil {
		return err
	}

	return nil
}

// GetAll News return all domain.news
func (r *NewsRepositoryImpl) GetAllByStatus(status string) ([]domain.News, error) {
	if status == "deleted" {
		news := []domain.News{}
		if err := r.Conn.Unscoped().Where("status = ?", status).Preload("Topic").Find(&news).Error; err != nil {
			return nil, err
		}

		return news, nil
	}

	news := []domain.News{}
	if err := r.Conn.Where("status = ?", status).Preload("Topic").Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}
