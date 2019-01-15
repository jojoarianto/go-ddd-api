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

	conn.AutoMigrate(&domain.News{})
	log.Println("Migration has been processed")
	// db.Model(&StockIns{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")

	return conn, nil
}
