package service

import (
	"log"
	"test-backend/internal/config"
	"test-backend/internal/model"
	"test-backend/internal/repository"

	"github.com/gofiber/fiber/v3"
)

type UserService struct {
	config         *config.Configuration
	userRepository repository.IUserRepository
}

type IUserService interface {
	GetUser(ctx fiber.Ctx, id string) (*model.User, error)
}

func NewUserService(
	config *config.Configuration,
	userRepository repository.IUserRepository,
) IUserService {
	return &UserService{
		config:         config,
		userRepository: userRepository,
	}
}

// GetUser...
func (s *UserService) GetUser(ctx fiber.Ctx, id string) (*model.User, error) {
	log.Printf("[Service:GetUser] UserId: %s", id)
	return s.userRepository.GetUser(ctx, id)
}
