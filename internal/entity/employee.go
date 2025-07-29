package entity

import "test-backend/internal/model"

type Employee struct {
	Id   int    `csv:"id" json:"id"`
	Name string `csv:"name" json:"name"`
	Age  int    `csv:"age" json:"age"`
	Team string `csv:"team" json:"team"`
}

type GetUserResponse struct {
	Status string      `json:"status"`
	Data   *model.User `json:"data"`
}
