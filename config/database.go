package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// ConnectDB returns initialized gorm.DB
func ConnectDB() (*gorm.DB, error) {
	user := getEnvWithDefault("DB_USER", "")
	password := getEnvWithDefault("DB_PASSWORD", "")
	host := getEnvWithDefault("DB_HOST", "")
	port := getEnvWithDefault("DB_PORT", "")
	dbname := getEnvWithDefault("DB_NAME", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// run and fill env with DB_PASSWORD="password" DB_NAME="db_name" go run main.go
func getEnvWithDefault(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}
