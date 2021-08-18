package repository

import (
	"garyshker"
	"github.com/jinzhu/gorm"
)

type Authorization interface {
	CreateUser(user *garyshker.User) (uint64, error)
	GetUser(usernameOrEmail, password string, isEmail bool) (*garyshker.User, error)
	FetchAuth(authD *garyshker.AuthDetails) (*garyshker.Auth, error)
	DeleteAuth(authD *garyshker.AuthDetails) error
	CreateAuth(userId uint64) (*garyshker.Auth, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
