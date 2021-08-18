package repository

import (
	"garyshker"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/twinj/uuid"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user *garyshker.User) (uint64, error) {
	role := "User"
	if user.Username == "admin" {
		role = "Admin"
	}
	user.Role = role
	err := a.db.Debug().Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (a *AuthPostgres) GetUser(usernameOrEmail, password string, isEmail bool) (*garyshker.User, error) {
	user := &garyshker.User{}
	var err error
	if isEmail {
		err = a.db.Debug().Where("email = ? AND password_hash = ?", usernameOrEmail, password).Take(&user).Error
	} else {
		err = a.db.Debug().Where("username = ? AND password_hash = ?", usernameOrEmail, password).Take(&user).Error
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

//func (s *Server) ValidateEmail(email string) error {
//	if email == "" {
//		return errors.New("required email")
//	}
//	if email != "" {
//		if err := checkmail.ValidateFormat(email); err != nil {
//			return errors.New("invalid email")
//		}
//	}
//	return nil
//}

func (a *AuthPostgres) FetchAuth(authD *garyshker.AuthDetails) (*garyshker.Auth, error) {
	au := &garyshker.Auth{}
	err := a.db.Debug().Where("user_id = ? AND auth_uuid = ?", authD.UserId, authD.AuthUuid).Take(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

func (a *AuthPostgres) DeleteAuth(authD *garyshker.AuthDetails) error {
	au := &garyshker.Auth{}
	db := a.db.Debug().Where("user_id = ? AND auth_uuid = ?", authD.UserId, authD.AuthUuid).Take(&au).Delete(&au)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

//Once the user signup/login, create a row in the auth table, with a new uuid
func (a *AuthPostgres) CreateAuth(userId uint64) (*garyshker.Auth, error) {
	au := &garyshker.Auth{}
	au.AuthUUID = uuid.NewV4().String() //generate a new UUID each time
	au.UserID = userId
	err := a.db.Debug().Create(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}