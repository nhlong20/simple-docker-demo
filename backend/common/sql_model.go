package common

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type SQLModel struct {
	Id        uuid.UUID  `json:"-" gorm:"column:id"`
	Status    int        `json:"status" gorm:"column:status;default:1;'"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (s *SQLModel) BeforeCreate(scope *gorm.DB) error {
	s.Id = uuid.New()
	return nil
}
