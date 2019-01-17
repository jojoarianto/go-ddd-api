package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
)

var config = ConfigDB{}

// ConfigDB db seting
type ConfigDB struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

// ConnectDB returns initialized gorm.DB
func ConnectDB() (*gorm.DB, error) {
	config.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Dbname)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Read and parse the configuration file
func (c *ConfigDB) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
