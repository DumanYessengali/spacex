package service

import (
	"errors"
	"garyshker"
	"garyshker/pkg/repository"
)

type CourseService struct {
	repos repository.Courses
}

func NewCourseService(repos repository.Courses) *CourseService {
	return &CourseService{repos: repos}
}

func (c *CourseService) GetAllCourse() (*[]garyshker.Course, error) {
	return c.repos.GetAllCourse()
}

func (c *CourseService) GetCourseById(courseId int) (*garyshker.Course, error) {
	return c.repos.GetCourseById(courseId)
}

func (c *CourseService) CreateCourse(course *garyshker.Course) (*garyshker.Course, error) {
	return c.repos.CreateCourse(course)
}

func (c *CourseService) UpdateCourse(courseName, courseDescription *string, course *garyshker.Course) (*garyshker.Course, error) {
	didUpdate := false
	if courseName != nil {
		course.CourseName = *courseName
		didUpdate = true
	}
	if courseDescription != nil {
		course.CourseDescription = *courseDescription
		didUpdate = true
	}
	if !didUpdate {
		return nil, errors.New("no update done")
	}

	return c.repos.UpdateCourse(course)
}

func (c *CourseService) UserCourseVerify(courseId int, userId uint64) (bool, error) {
	return c.repos.UserCourseVerify(courseId, userId)
}

func (c *CourseService) EnrollCourse(courseId int, userId uint64) error {
	return c.repos.EnrollCourse(courseId, userId)
}

func (c *CourseService) GetAllMyCourse(userId uint64) (*[]garyshker.Course, error) {
	return c.repos.GetAllMyCourse(userId)
}

func (c *CourseService) DeleteMyCourse(courseId int, userId uint64) error {
	return c.repos.DeleteMyCourse(courseId, userId)
}
