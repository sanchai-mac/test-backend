package service

import (
	"log"
	"test-backend/internal/config"
	"test-backend/internal/entity"
	"test-backend/internal/model"
	"test-backend/internal/repository"
	"test-backend/internal/util"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	config         *config.Configuration
	userRepository repository.IUserRepository
}

type IUserService interface {
	GetUser(ctx *gin.Context, id string) (*model.User, error)
	CreateUser(ctx *gin.Context, input *entity.CreateUserRequest) error
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
func (s *UserService) GetUser(ctx *gin.Context, id string) (*model.User, error) {
	log.Printf("[Service:GetUser] UserId: %s", id)
	return s.userRepository.GetUser(ctx, id)
}

// CreateUser...
func (s *UserService) CreateUser(ctx *gin.Context, input *entity.CreateUserRequest) error {
	log.Printf("[Service:GetUser] Request: %s", util.ConvertStructToJSONString(input))
	if err := s.userRepository.CreateUser(ctx, &model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}); err != nil {
		return err
	}
	return nil
}
