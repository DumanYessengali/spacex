package service

import (
	"garyshker"
	"garyshker/pkg/repository"
	"net/http"
)

type Authorization interface {
	CreateUser(user *garyshker.User) (uint64, error)
	GetUser(usernameOrEmail, password string) (*garyshker.User, error)
	FetchAuth(authD *garyshker.AuthDetails) (*garyshker.Auth, error)
	DeleteAuth(authD *garyshker.AuthDetails) error
	CreateAuth(userId uint64) (*garyshker.Auth, error)
	ExtractTokenAuth(r *http.Request) (*garyshker.AuthDetails, error)
	SignIn(authD garyshker.AuthDetails) (string, error)
	TokenValid(r *http.Request) error
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}
}
