package repository

import (
	"log"
	"test-backend/internal/config"
	"test-backend/internal/infrastructure/database"
	"test-backend/internal/model"

	"github.com/gofiber/fiber/v3"
)

type UserRepository struct {
	config *config.Configuration
	db     *database.DB
}

type IUserRepository interface {
	GetUser(ctx fiber.Ctx, id string) (*model.User, error)
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
func (r *UserRepository) GetUser(ctx fiber.Ctx, id string) (*model.User, error) {
	log.Printf("[Repository:GetUser] UserId: %s", id)

	user := &model.User{}
	if tx := r.db.CostomerDB.Where(`user_id = ?`, id).Find(&user); tx.Error != nil {
		log.Printf("[Repository:GetUser] Find user error: %s", tx.Error)
		return nil, tx.Error
	}
	return user, nil
}
