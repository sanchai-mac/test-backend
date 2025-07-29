package controller

import (
	"log"
	"net/http"
	"test-backend/internal/config"
	"test-backend/internal/entity"
	"test-backend/internal/model"
	"test-backend/internal/service"
	"test-backend/internal/util/error_wrapper"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	config       *config.Configuration
	iUserService service.IUserService
}

type IUserController interface {
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
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
// GetUser...
func (c *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("user_id")
	log.Println("[Controller:GetUser] Request user_id: ", id)
	user, err := c.iUserService.GetUser(ctx, id)
	if err != nil {
		if err.Error() == error_wrapper.NOT_FOUND.String() {
			ctx.JSON(http.StatusNotFound, entity.GetUserResponse{
				Status: "Not Found",
				Data:   &model.User{},
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, entity.GetUserResponse{
				Status: "Internal Server Error",
				Data:   &model.User{},
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, entity.GetUserResponse{
		Status: "Success",
		Data:   user,
	})
}

// CreateUser...
func (c *UserController) CreateUser(ctx *gin.Context) {
	log.Println("[Controller:CreateUser] Called...")

	input := &entity.CreateUserRequest{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.GetUserResponse{
			Status: "Bad Request",
		})
		return
	}

	err := c.iUserService.CreateUser(ctx, input)
	if err != nil {
		log.Println("[CreateUser] Service error:", err)
		ctx.JSON(http.StatusInternalServerError, entity.GetUserResponse{
			Status: "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, entity.GetUserResponse{
		Status: "Success",
	})
}

func (c *UserController) UpdateUser(ctx *gin.Context)
func (c *UserController) DeleteUser(ctx *gin.Context)
