package entity

import "test-backend/internal/model"

type GetUserResponse struct {
	Status string      `json:"status"`
	Data   *model.User `json:"data"`
}
