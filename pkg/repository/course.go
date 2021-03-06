package repository

import (
	"garyshker"
	"github.com/jinzhu/gorm"
)

type CoursePostgres struct {
	db *gorm.DB
}

func NewCoursePostgres(db *gorm.DB) *CoursePostgres {
	return &CoursePostgres{db: db}
}

func (c *CoursePostgres) GetAllCourse() (*[]garyshker.Course, error) {
	courses := &[]garyshker.Course{}
	err := c.db.Debug().Find(&courses).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CoursePostgres) GetCourseById(courseId int) (*garyshker.Course, error) {
	course := &garyshker.Course{}
	err := c.db.Debug().Where("id = ?", courseId).Take(&course).Error
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (c *CoursePostgres) CreateCourse(course *garyshker.Course) (*garyshker.Course, error) {
	err := c.db.Debug().Create(&course).Error
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c *CoursePostgres) UpdateCourse(course *garyshker.Course) (*garyshker.Course, error) {
	err := c.db.Debug().Model(&course).Updates(garyshker.Course{
		CourseName:        course.CourseName,
		CourseDescription: course.CourseDescription,
	}).Error
	if err != nil {
		return nil, err
	}

	return course, nil
}

var result struct {
	Found bool
}

func (c *CoursePostgres) UserCourseVerify(courseId int, userId uint64) (bool, error) {
	err := c.db.Raw("SELECT EXISTS(SELECT 1 FROM user_courses WHERE course_id = ? and user_id = ?) AS found", courseId, userId).Scan(&result).Error
	if err != nil {
		return false, err
	}
	return result.Found, nil
}

func (c *CoursePostgres) EnrollCourse(courseId int, userId uint64) error {
	userCourse := &garyshker.UserCourse{}
	userCourse.CourseId = uint64(courseId)
	userCourse.UserId = userId

	err := c.db.Debug().Create(&userCourse).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CoursePostgres) GetAllMyCourse(userId uint64) (*[]garyshker.Course, error) {
	courses := &[]garyshker.Course{}

	err := c.db.Debug().Raw("select * from courses where id in (select course_id from user_courses where user_id = $1)", userId).Scan(&courses).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CoursePostgres) DeleteMyCourse(courseId int, userId uint64) error {
	userCourse := &garyshker.UserCourse{}
	err := c.db.Debug().Where("course_id = ? and user_id = ?", courseId, userId).Delete(&userCourse).Error
	if err != nil {
		return err
	}
	return nil
}
