package service

import (
	"log"
	"test-backend/internal/config"
	"test-backend/internal/model"

	"github.com/gofiber/fiber/v3"
)

type UserService struct {
	config *config.Configuration
}

type IUserService interface {
	GetUser(ctx fiber.Ctx, id string) (*model.User, error)
}

func NewUserService(
	config *config.Configuration,
) IUserService {
	return &UserService{
		config: config,
	}
}

// GetUser...
func (s *UserService) GetUser(ctx fiber.Ctx, id string) (*model.User, error) {
	log.Printf("[Service:GetUser] UserId: %s", id)

	return &model.User{}, nil
}
