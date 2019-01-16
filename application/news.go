package application

import (
	"github.com/jojoarianto/go-ddd-api/config"
	"github.com/jojoarianto/go-ddd-api/domain"
	"github.com/jojoarianto/go-ddd-api/infrastructure/persistence"
)

// GetNews returns domain.news by id
func GetNews(id int) (*domain.News, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.Get(id)
}

// GetAllNews return all domain.news
func GetAllNews() ([]domain.News, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.GetAll()
}

// AddNews saves new news
func AddNews(p domain.News) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.Save(&p)
}

// RemoveNews do remove news by id
func RemoveNews(id int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.Remove(id)
}

// UpdateNews do remove news by id
func UpdateNews(p domain.News, id int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	p.ID = uint(id)

	return repo.Update(&p)
}

// GetAllNewsByFilter return all domain.news by filter
func GetAllNewsByFilter(status string) ([]domain.News, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.GetAllByStatus(status)
}

// GetNewsByTopic returns []domain.news by topic.slug
func GetNewsByTopic(slug string) ([]*domain.News, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.GetBySlug(slug)
}
