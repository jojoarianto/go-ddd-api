package persistence

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jojoarianto/go-ddd-api/domain"
	"github.com/jojoarianto/go-ddd-api/domain/repository"
)

// TopicRepositoryImpl Implements repository.TopicRepository
type TopicRepositoryImpl struct {
	Conn *gorm.DB
}

// NewTopicRepositoryWithRDB returns initialized TopicRepositoryImpl
func NewTopicRepositoryWithRDB(conn *gorm.DB) repository.TopicRepository {
	return &TopicRepositoryImpl{Conn: conn}
}

// Get topic by id return domain.topic
func (r *TopicRepositoryImpl) Get(id int) (*domain.Topic, error) {
	topic := &domain.Topic{}
	if err := r.Conn.Preload("News").First(&topic, id).Error; err != nil {
		return nil, err
	}
	return topic, nil
}

// GetAll topic return all domain.topic
func (r *TopicRepositoryImpl) GetAll() (*domain.Topic, error) {
	// topic := &domain.Topic{}
	// if err := r.Conn.First(&topic, id).Error; err != nil {
	// 	return nil, err
	// }
	// return topic, nil
	return nil, nil
}

// Save return
func (r *TopicRepositoryImpl) Save(topic *domain.Topic) error {
	return nil
}

func (r *TopicRepositoryImpl) Remove(topic *domain.Topic) error {
	return nil
}
