package garyshker

//type User struct {
//	Id       uint64 `gorm:"primary_key;auto_increment" json:"-" db:"id"`
//	Name     string `gorm:"size:255;not null;" json:"name" binding:"required,min=2,max=25"`
//	Username string `gorm:"size:255;not null;unique;" json:"username" binding:"required,min=2,max=25"`
//	Password string `gorm:"size:255;not null;" json:"password" binding:"required,min=4"`
//	Email    string `gorm:"size:255;not null;unique;" json:"email" binding:"required,email"`
//	Role     string `gorm:";not null;" json:"-" db:"role"`
//}
//
//type Auth struct {
//	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
//	UserID   uint64 `gorm:";not null;" json:"user_id"`
//	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
//}

type User struct {
	Id       uint64 `json:"-" db:"id"`
	Name     string `json:"name" binding:"required,min=2,max=25"`
	Username string `json:"username" binding:"required,min=2,max=25"`
	Password string `json:"password" binding:"required,min=4"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"-" db:"role"`
}

type Auth struct {
	ID       uint64 `json:"id"`
	UserID   uint64 `json:"user_id"`
	AuthUUID string `json:"auth_uuid"`
}

type AuthDetails struct {
	AuthUuid string
	UserId   uint64
}
