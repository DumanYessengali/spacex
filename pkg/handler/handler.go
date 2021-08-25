package handler

import (
	"garyshker/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.CreateUser)
		auth.POST("/sign-in", h.Login)
	}

	api := router.Group("api", h.TokenAuthMiddleware())
	{
		api.POST("/logout", h.LogOut)
		api.POST("/enroll-course/:id", h.enrollCourse)
		api.POST("/enroll-post/:id", h.enrollPost)

		profile := api.Group("/profile")
		{
			profile.GET("/", h.getUser)
			profile.PUT("/", h.updateUser)
		}

		myCourse := api.Group("/my-courses")
		{
			myCourse.GET("/", h.getAllMyCourse)
			myCourse.DELETE("/:id", h.deleteMyCourse)
		}

		mySavedPost := api.Group("/my-saved-post")
		{
			mySavedPost.GET("/", h.getAllMySavedPost)
			mySavedPost.DELETE("/:id", h.deleteMySavedPost)
		}

		admin := api.Group("/admin", h.requireAdmin())
		{
			course := admin.Group("/course")
			{
				course.POST("/create-course", h.createCourse)
				course.PUT("/update-course/:id", h.updateCourse)
			}

			post := admin.Group("/post")
			{
				post.POST("/create-video-post", h.createVideoPost)
				post.POST("/create-article-post", h.createArticlePost)
				post.PUT("/update-post/:id", h.updatePost)
			}
		}
	}

	course := router.Group("/course")
	{
		course.GET("/", h.getAllCourse)
		course.GET("/:id", h.getCourseById)
	}

	post := router.Group("/post")
	{
		post.GET("/", h.getAllPost)
		post.GET("/:id", h.getPostById)
	}

	return router
}
