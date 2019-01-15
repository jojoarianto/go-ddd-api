package application

import (
	"github.com/jojoarianto/go-ddd-api/config"
	"github.com/jojoarianto/go-ddd-api/domain"
	"github.com/jojoarianto/go-ddd-api/infrastructure/persistence"
)

// GetUser returns user
func GetTopic(id int) (*domain.Topic, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewTopicRepositoryWithRDB(conn)
	return repo.Get(id)
}
