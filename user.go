package garyshker

type User struct {
	Id       uint64 `json:"-" db:"id"`
	Name     string `json:"name" binding:"required,min=2,max=25"`
	Username string `json:"username" binding:"required,min=2,max=25"`
	Password string `json:"password" binding:"required,min=4"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"-" db:"role"`
}

type Auth struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID   uint64 `gorm:";not null;" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
}

type AuthDetails struct {
	AuthUuid string
	UserId   uint64
}
