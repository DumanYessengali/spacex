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

type Users interface {
	GetUserByUserId(userId uint64) (*garyshker.UserAllInformation, error)
	GetUserInfo(userId uint64) (*garyshker.UserInformation, error)
	GetUser(userId uint64) (*garyshker.User, error)
	UpdateUser(userInfo *garyshker.UserInformation, user *garyshker.User) (*garyshker.UserAllInformation, error)
	GetRole(id uint64) (garyshker.Role, error)
}

type Courses interface {
	GetAllCourse() (*[]garyshker.Course, error)
	GetCourseById(courseId int) (*garyshker.Course, error)
	CreateCourse(course *garyshker.Course) (*garyshker.Course, error)
	UpdateCourse(course *garyshker.Course) (*garyshker.Course, error)
	EnrollCourse(courseId int, userId uint64) error
	GetAllMyCourse(userId uint64) (*[]garyshker.Course, error)
	DeleteMyCourse(courseId int, userId uint64) error
	UserCourseVerify(courseId int, userId uint64) (bool, error)
}

type Posts interface {
	GetAllPost() (*[]garyshker.VideoPost, *[]garyshker.ArticlePost, error)
	GetPostById(postId int) (interface{}, *garyshker.PostConnection, error)
	GetVideoPostById(id uint64) (*garyshker.VideoPost, error)
	GetArticlePostById(id uint64) (*garyshker.ArticlePost, error)
	EnrollPost(post *garyshker.PostConnection, userId uint64) error
	GetAllMySavedPosts(userId uint64) (*[]garyshker.VideoPost, *[]garyshker.ArticlePost, error)
	CreateVideoPost(videoPost *garyshker.VideoPost) (*garyshker.VideoPost, error)
	CreateArticlePost(articlePost *garyshker.ArticlePost) (*garyshker.ArticlePost, error)
	DeleteMySavedPost(postId, userId uint64) error
	UpdateVideoPost(videoPost *garyshker.VideoPost) (*garyshker.VideoPost, error)
	UpdateArticlePost(articlePost *garyshker.ArticlePost) (*garyshker.ArticlePost, error)
	UserPostVerify(post *garyshker.PostConnection, userId uint64) (bool, error)
}

type Repository struct {
	Authorization
	Users
	Courses
	Posts
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Users:         NewUserPostgres(db),
		Courses:       NewCoursePostgres(db),
		Posts:         NewPostPostgres(db),
	}
}
