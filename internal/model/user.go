package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId    uuid.UUID  `json:"user_id"`
	UserName  string     `json:"user_name"`
	LastName  string     `json:"last_name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
