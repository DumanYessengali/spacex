package service

import (
	"errors"
	"garyshker"
	"garyshker/pkg/repository"
)

type UserService struct {
	repos repository.Users
}

func NewUserService(repos repository.Users) *UserService {
	return &UserService{repos: repos}
}
func (u *UserService) GetUserByUserId(userId uint64) (*garyshker.UserAllInformation, error) {
	return u.repos.GetUserByUserId(userId)
}
func (u *UserService) GetUserInfo(userId uint64) (*garyshker.UserInformation, error) {
	return u.repos.GetUserInfo(userId)
}

func (u *UserService) UpdateUser(userAge *int, userAvatarImage, userCity *string, userInfo *garyshker.UserInformation, username, name, email *string, user *garyshker.User) (*garyshker.UserAllInformation, error) {
	didUpdate := false

	if userAge != nil {
		userInfo.UserAge = *userAge
		didUpdate = true
	}

	if userAvatarImage != nil {
		userInfo.UserAvatarImage = *userAvatarImage
		didUpdate = true
	}

	if userCity != nil {
		userInfo.UserCity = *userCity
		didUpdate = true
	}

	if username != nil {
		user.Username = *username
		didUpdate = true
	}

	if name != nil {
		user.Name = *name
		didUpdate = true
	}

	if email != nil {
		user.Email = *email
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	return u.repos.UpdateUser(userInfo, user)
}
func (u *UserService) GetUser(userId uint64) (*garyshker.User, error) {
	return u.repos.GetUser(userId)

}
