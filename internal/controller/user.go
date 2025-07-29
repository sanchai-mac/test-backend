package controller

import (
	"log"
	"test-backend/internal/config"
	"test-backend/internal/entity"
	"test-backend/internal/model"
	"test-backend/internal/service"
	"test-backend/internal/util/error_wrapper"

	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	config       *config.Configuration
	iUserService service.IUserService
}

type IUserController interface {
	GetUser(ctx fiber.Ctx) error
}

func NewUserController(
	config *config.Configuration,
	iUserService service.IUserService,
) IUserController {
	return &UserController{
		config:       config,
		iUserService: iUserService,
	}
}

// GetUser...
func (c *UserController) GetUser(ctx fiber.Ctx) error {
	id := ctx.Params("user_id")
	log.Println("[Controller:GetUser] Request user_id: ", id)
	user, err := c.iUserService.GetUser(ctx, id)
	if err != nil {
		if err.Error() == error_wrapper.NOT_FOUND.String() {
			return ctx.Status(fiber.StatusNotFound).JSON(entity.GetUserResponse{
				Status: "Not Found",
				Data:   &model.User{},
			})
		} else {
			return ctx.Status(fiber.StatusInternalServerError).JSON(entity.GetUserResponse{
				Status: "Internal Server Error",
				Data:   &model.User{},
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(entity.GetUserResponse{
		Status: "Success",
		Data:   user,
	})
}
