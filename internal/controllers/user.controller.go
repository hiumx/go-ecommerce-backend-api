package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiumx/go-ecommerce-backend-api/internal/services"
	"github.com/hiumx/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		response.ErrorResponse(c, 20002)
		return
	}
	response.SuccessResponse(c, 20001, uc.userService.GetUserByIdService(id))
}
