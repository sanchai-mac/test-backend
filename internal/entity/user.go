package entity

import "test-backend/internal/model"

type GetUserResponse struct {
	Status string      `json:"status"`
	Data   *model.User `json:"data"`
}

type UserRequest struct {
	UserId    string `json:"user_id" `
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
}
