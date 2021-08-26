package repository

import (
	"garyshker"
	"github.com/jinzhu/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}
func (u *UserPostgres) GetUserByUserId(userId uint64) (*garyshker.UserAllInformation, error) {
	user := &garyshker.User{}
	userInfo := &garyshker.UserInformation{}
	err := u.db.Debug().Where("id = ?", userId).Take(&user).Error
	if err != nil {
		return nil, err
	}
	err = u.db.Debug().Where("user_id = ?", userId).Take(&userInfo).Error
	if err != nil {
		return nil, err
	}
	userAllInfo := &garyshker.UserAllInformation{
		UserAvatarImage: userInfo.UserAvatarImage,
		UserCity:        userInfo.UserCity,
		UserAge:         userInfo.UserAge,
		UserId:          userId,
		Name:            user.Name,
		Username:        user.Username,
		Email:           user.Email,
	}
	return userAllInfo, nil
}

func (u *UserPostgres) UpdateUser(userInfo *garyshker.UserInformation, user *garyshker.User) (*garyshker.UserAllInformation, error) {
	err := u.db.Debug().Model(&userInfo).Updates(garyshker.UserInformation{
		UserAvatarImage: userInfo.UserAvatarImage,
		UserCity:        userInfo.UserCity,
		UserAge:         userInfo.UserAge,
	}).Error

	if err != nil {
		return nil, err
	}
	err = u.db.Debug().Model(&user).Updates(garyshker.User{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}).Error

	if err != nil {
		return nil, err
	}
	userAllInfo := &garyshker.UserAllInformation{
		UserAvatarImage: userInfo.UserAvatarImage,
		UserCity:        userInfo.UserCity,
		UserAge:         userInfo.UserAge,
		UserId:          user.Id,
		Name:            user.Name,
		Username:        user.Username,
		Email:           user.Email,
	}

	return userAllInfo, nil
}

func (u *UserPostgres) GetUserInfo(userId uint64) (*garyshker.UserInformation, error) {
	userInfo := &garyshker.UserInformation{}
	err := u.db.Debug().Where("user_id = ?", userId).Take(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (u *UserPostgres) GetUser(userId uint64) (*garyshker.User, error) {
	user := &garyshker.User{}
	err := u.db.Debug().Where("id = ?", userId).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserPostgres) GetRole(id uint64) (garyshker.Role, error) {
	user := &garyshker.User{}
	err := u.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return "", err
	}
	return user.Role, nil
}
