package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId    uuid.UUID  `json:"user_id" gorm:"column:user_id;type:uuid;primaryKey"`
	FirstName string     `json:"first_name" gorm:"column:first_name"`
	LastName  string     `json:"last_name" gorm:"column:last_name"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
