package adapters

import (
	"fmt"
	"garyshker"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

//func NewPostgresDB(cfg Config) (*gorm.DB, error) {
//	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
//		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
//	if err != nil {
//		return nil, err
//	}
//
//	err = db.Ping()
//	if err != nil {
//		return nil, err
//	}
//
//	return db, nil
//}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode, cfg.Password)
	db, err := gorm.Open("postgres", DBURL)
	if err != nil {
		return nil, err
	}
	db.Debug().AutoMigrate(
		&garyshker.Auth{},
		&garyshker.User{},
	)
	return db, nil
}
