package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hiumx/go-ecommerce-backend-api/internal/services"
	"github.com/hiumx/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(
	userService services.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	response.SuccessResponse(c, 20001, uc.userService.Register("abc@gmail.com", "register"))
}
