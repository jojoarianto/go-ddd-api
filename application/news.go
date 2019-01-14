package application

import (
	"github.com/jojoarianto/go-ddd-api/config"
	"github.com/jojoarianto/go-ddd-api/domain"
	"github.com/jojoarianto/go-ddd-api/infrastructure/persistence"
)

// GetUser returns user
func GetNews(id int) (*domain.News, error) {
	conn, err := config.NewDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.Get(id)
}
