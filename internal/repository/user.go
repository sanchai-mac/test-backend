package repository

import (
	"log"
	"test-backend/internal/config"
	"test-backend/internal/infrastructure/database"
	"test-backend/internal/model"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	config *config.Configuration
	db     *database.DB
}

type IUserRepository interface {
	GetUser(ctx *gin.Context, id string) (*model.User, error)
	CreateUser(ctx *gin.Context, input *model.User) error
}

func NewUserRepository(
	config *config.Configuration,
	db *database.DB,
) IUserRepository {
	return &UserRepository{
		config: config,
		db:     db,
	}
}

// GetUser...
func (r *UserRepository) GetUser(ctx *gin.Context, id string) (*model.User, error) {
	log.Printf("[Repository:GetUser] UserId: %s", id)

	user := &model.User{}
	if tx := r.db.CostomerDB.Where(`user_id = ?`, id).Find(&user); tx.Error != nil {
		log.Printf("[Repository:GetUser] Find user error: %s", tx.Error)
		return nil, tx.Error
	}
	return user, nil
}

// CreateUser...
func (r *UserRepository) CreateUser(ctx *gin.Context, input *model.User) error {
	log.Printf("[Repository:GetUser] Called...")

	if tx := r.db.CostomerDB.Create(&input); tx.Error != nil {
		log.Printf("[Repository:GetUser] Create user error: %s", tx.Error)
		return tx.Error
	}
	return nil
}
