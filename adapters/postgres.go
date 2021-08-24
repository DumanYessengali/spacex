package adapters

import (
	"fmt"
	"garyshker"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode, cfg.Password)
	db, err := gorm.Open("postgres", DBURL)
	if err != nil {
		return nil, err
	}
	db.Debug().AutoMigrate(
		&garyshker.Auth{},
		&garyshker.User{},
		&garyshker.UserInformation{},
		&garyshker.Course{},
	)
	return db, nil
}
