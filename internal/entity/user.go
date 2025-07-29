package entity

import "test-backend/internal/model"

type GetUserResponse struct {
	Status string      `json:"status"`
	Data   *model.User `json:"data"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
}
