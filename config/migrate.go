package config

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/go-ddd-api/domain"
)

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func DBMigrate() (*gorm.DB, error) {
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.AutoMigrate(domain.News{}, domain.Topic{})
	log.Println("Migration has been processed")

	// conn.Model(&domain.News{}).Related(&domain.Topic{})

	// Example
	// db.Model(&language).Related(&users)
	// SELECT * FROM "users" INNER JOIN "user_languages" ON "user_languages"."user_id" = "users"."id" WHERE  ("user_languages"."language_id" IN ('111'))

	return conn, nil
}
