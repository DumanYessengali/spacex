package garyshker

import "time"

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
	Role     Role   `json:"-" db:"role"`
}

type Role string

const (
	AdminRole Role = "Admin"
	UserRole  Role = "User"
)

type Auth struct {
	ID       uint64 `json:"id"`
	UserID   uint64 `json:"user_id"`
	AuthUUID string `json:"auth_uuid"`
}

type AuthDetails struct {
	AuthUuid string
	UserId   uint64
}

type UserInformation struct {
	Id              uint64 `json:"-" db:"id"`
	UserAvatarImage string `json:"user_avatar_image"`
	UserCity        string `json:"user_city"`
	UserAge         int    `json:"user_age"`
	UserId          uint64 `json:"-" db:"user_id"`
}

type UserAllInformation struct {
	UserId          uint64
	Name            string
	Username        string
	Email           string
	UserAvatarImage string
	UserCity        string
	UserAge         int
}

type Course struct {
	Id                uint64 `json:"-" db:"id"`
	CourseName        string `json:"course_name"`
	CourseDescription string `json:"course_description"`
}

type UserCourse struct {
	Id       uint64
	CourseId uint64
	UserId   uint64
}

type VideoPost struct {
	Id            uint64        `json:"-" db:"id"`
	Title         string        `json:"title"`
	TitleType     PostTitleType `json:"title_type"`
	VideoDuration int           `json:"video_duration"`
	Description   string        `json:"description"`
	VideoUrl      string        `json:"video_url"`
	Created       time.Time     `json:"-" db:"created"`
	Updated       time.Time     `json:"-" db:"updated"`
}

type ArticlePost struct {
	Id                         uint64        `json:"-" db:"id"`
	Title                      string        `json:"title"`
	TitleType                  PostTitleType `json:"title_type"`
	Duration                   int           `json:"duration"`
	AuthorInformationParagraph string        `json:"author_information_paragraph"`
	ParagraphName              string        `json:"paragraph_name"`
	Description                string        `json:"description"`
	AuthorName                 string        `json:"author_name"`
	AuthorPosition             string        `json:"author_position"`
	Created                    time.Time     `json:"-" db:"created"`
	Updated                    time.Time     `json:"-" db:"updated"`
}

type PostConnection struct {
	Id       uint64
	PostId   uint64
	PostType PostType
}

type PostType string

const (
	VideoPosts   PostType = "video"
	ArticlePosts PostType = "article"
)

type PostTitleType string

const (
	EcoPhilosophyPost     PostTitleType = "Экофилософия"
	FinancialLiteracyPost PostTitleType = "Фин.грамотность"
	MentalHealthPost      PostTitleType = "Ментальное здоровье"
	CreationPost          PostTitleType = "Творчество"
	SexEducationPost      PostTitleType = "Половое воспитание"
	CulturePost           PostTitleType = "Культура"
)

type UserSavedPost struct {
	Id               uint64
	PostConnectionId uint64
	UserId           uint64
}
