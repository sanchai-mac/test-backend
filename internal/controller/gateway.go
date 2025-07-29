package controller

import (
	"go.uber.org/dig"
)

// Controller ...
type Controller struct {
	Gateway Gateway
}

// Gateway ...
type Gateway struct {
	dig.In
	IUserController IUserController
}
