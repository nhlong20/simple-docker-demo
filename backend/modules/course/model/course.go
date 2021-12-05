package coursemodel

import (
	"btcn03-api/common"
	"errors"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

const EntityName = "Course"

type Course struct {
	common.SQLModel
	Code         string `json:"code" gorm:"column:code;unique;not null;varchar;size:20"`
	Slug         string `json:"slug" gorm:"column:slug;unique;not null;"`
	Title        string `json:"title" gorm:"column:title; not null;"`
	Department   string `json:"department" gorm:"column:department;not null;varchar;size:200"`
	TeacherName  string `json:"teacher_name" gorm:"column:teacher_name;varchar;size:256"`
	Description  string `json:"description" gorm:"column:description;"`
	AcademicYear string `json:"acad_year" gorm:"column:acad_year;varchar;size:10;"`
	Semester     string `json:"semester" gorm:"column:semester;type:json"`
	Credit       string `json:"credit" gorm:"column:credit;default:'2'"`
}

var (
	ErrTitleCannotBeBlank      = errors.New("Title cannot be blank")
	ErrCourseCodeCannotBeBlank = errors.New("Course code cannot be blank")
	ErrDepartmentCannotBeBlank = errors.New("Department code cannot be blank")
)

func (Course) TableName() string { return "courses" }

func (c *Course) BeforeSave(scope *gorm.DB) error {
	c.Slug = slug.Make(c.Title) + "-" + strconv.FormatInt(time.Now().UnixNano(), 10)
	return nil
}
func (data *Course) Validate() error {
	data.Title = strings.TrimSpace(data.Title)

	if data.Title == "" {
		return ErrTitleCannotBeBlank
	}

	data.Department = strings.TrimSpace(data.Department)

	if data.Department == "" {
		return ErrDepartmentCannotBeBlank
	}

	data.Code = strings.TrimSpace(data.Code)

	if data.Title == "" {
		return ErrCourseCodeCannotBeBlank
	}

	return nil
}
