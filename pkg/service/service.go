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
type Users interface {
	GetUserByUserId(userId uint64) (*garyshker.UserAllInformation, error)
	GetUserInfo(userId uint64) (*garyshker.UserInformation, error)
	GetUser(userId uint64) (*garyshker.User, error)
	UpdateUser(userAge *int, userAvatarImage, userCity *string, userInfo *garyshker.UserInformation, username, name, email *string, user *garyshker.User) (*garyshker.UserAllInformation, error)
}

type Courses interface {
	GetAllCourse() (*[]garyshker.Course, error)
	GetCourseById(courseId int) (*garyshker.Course, error)
	CreateCourse(course *garyshker.Course) (*garyshker.Course, error)
	UpdateCourse(courseName, courseDescription *string, course *garyshker.Course) (*garyshker.Course, error)
	EnrollCourse(courseId int, userId uint64) error
	GetAllMyCourse(userId uint64) (*[]garyshker.Course, error)
	DeleteMyCourse(courseId int, userId uint64) error
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
	UpdateVideoPost(title, description, videoUrl *string, titleType *garyshker.PostTitleType, videoDuration *int, videoPost *garyshker.VideoPost) (*garyshker.VideoPost, error)
	UpdateArticlePost(title, authorInformationParagraph, paragraphName, description, authorName, authorPosition *string, duration *int, titleType *garyshker.PostTitleType, articlePost *garyshker.ArticlePost) (*garyshker.ArticlePost, error)
}

type Service struct {
	Authorization
	Users
	Courses
	Posts
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Users:         NewUserService(repository.Users),
		Courses:       NewCourseService(repository.Courses),
		Posts:         NewPostService(repository.Posts),
	}
}
