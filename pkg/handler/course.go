package handler

import (
	"garyshker"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllCourse(c *gin.Context) {
	courses, err := h.services.Courses.GetAllCourse()
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (h *Handler) getCourseById(c *gin.Context) {
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	course, err := h.services.Courses.GetCourseById(courseId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *Handler) getAllMyCourse(c *gin.Context) {
	token, err := h.services.Authorization.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	foundAuth, err := h.services.Authorization.FetchAuth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	courses, err := h.services.Courses.GetAllMyCourse(foundAuth.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) createCourse(c *gin.Context) {
	token, err := h.services.Authorization.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	_, err = h.services.Authorization.FetchAuth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var course garyshker.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, "invalid json")
		return
	}

	courseTable, err := h.services.Courses.CreateCourse(&course)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, courseTable)
}

type UpdateCourse struct {
	CourseName        *string `json:"course_name"`
	CourseDescription *string `json:"course_description"`
}

func (h *Handler) updateCourse(c *gin.Context) {
	token, err := h.services.Authorization.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	_, err = h.services.Authorization.FetchAuth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	course, err := h.services.Courses.GetCourseById(courseId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var u UpdateCourse

	if err := c.ShouldBindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	updateCourseTable, err := h.services.Courses.UpdateCourse(u.CourseName, u.CourseDescription, course)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateCourseTable)
}

func (h *Handler) enrollCourse(c *gin.Context) {
	token, err := h.services.Authorization.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	foundAuth, err := h.services.Authorization.FetchAuth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	course, err := h.services.Courses.GetCourseById(courseId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	checkUserInCourse, err := h.services.Courses.UserCourseVerify(courseId, foundAuth.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	if checkUserInCourse {
		c.JSON(http.StatusOK, "You are already studying in this course "+course.CourseName)
	} else {
		err = h.services.Courses.EnrollCourse(courseId, foundAuth.UserID)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusOK, "You are successfully joined to the course "+course.CourseName)
	}
}

func (h *Handler) deleteMyCourse(c *gin.Context) {
	token, err := h.services.Authorization.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	foundAuth, err := h.services.Authorization.FetchAuth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	course, err := h.services.Courses.GetCourseById(courseId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	err = h.services.Courses.DeleteMyCourse(courseId, foundAuth.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, "You are successfully remove from course "+course.CourseName)
}
