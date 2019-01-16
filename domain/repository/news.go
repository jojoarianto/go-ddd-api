package repository

import "github.com/jojoarianto/go-ddd-api/domain"

// NewsRepository represent repository of  the news
// Expect implementation by the infrastructure layer
type NewsRepository interface {
	Get(id int) (*domain.News, error)
	GetAll() ([]domain.News, error)
	GetAllByStatus(status string) ([]domain.News, error)
	Save(*domain.News) error
	Remove(id int) error
	Update(*domain.News) error
}
