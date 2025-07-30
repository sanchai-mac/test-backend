package repository

import (
	"errors"
	"log"
	"test-backend/internal/config"
	"test-backend/internal/infrastructure/database"
	"test-backend/internal/model"
	"test-backend/internal/util/error_wrapper"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	config *config.Configuration
	db     *database.DB
}

type IUserRepository interface {
	GetUser(ctx *gin.Context, id string) (*model.User, error)
	CreateUser(ctx *gin.Context, input *model.User) error
	UpdateUser(ctx *gin.Context, input *model.User) error
	DeleteUserByUserId(ctx *gin.Context, input *model.User) error
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
	tx := r.db.CostomerDB.Where(`user_id = ?`, id).Find(&user)
	if tx.Error != nil {
		log.Printf("[Repository:GetUser] Find user error: %s", tx.Error)
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		log.Printf("[Repository:GetUser] No user found with id: %s", id)
		return nil, errors.New(error_wrapper.NOT_FOUND.String())
	}
	return user, nil
}

// CreateUser...
func (r *UserRepository) CreateUser(ctx *gin.Context, input *model.User) error {
	log.Printf("[Repository:CreateUser] Called...")

	if tx := r.db.CostomerDB.Create(&input); tx.Error != nil {
		log.Printf("[Repository:CreateUser] Create user error: %s", tx.Error)
		return tx.Error
	}
	return nil
}

// UpdateUser...
func (r *UserRepository) UpdateUser(ctx *gin.Context, input *model.User) error {
	log.Printf("[Repository:UpdateUser] UserId: ", input)

	if tx := r.db.CostomerDB.Where(`user_id = ?`, input.UserId).Updates(&input); tx.Error != nil {
		log.Printf("[Repository:UpdateUser] Updates user error: %s", tx.Error)
		return tx.Error
	}
	return nil
}

// DeleteUserByUserId...
func (r *UserRepository) DeleteUserByUserId(ctx *gin.Context, input *model.User) error {
	log.Printf("[Repository:DeleteUserByUserId] UserId: ", input)

	if tx := r.db.CostomerDB.Where(`user_id = ?`, input.UserId).Delete(&input); tx.Error != nil {
		log.Printf("[Repository:DeleteUserByUserId] Delete user error: %s", tx.Error)
		return tx.Error
	}
	return nil
}
