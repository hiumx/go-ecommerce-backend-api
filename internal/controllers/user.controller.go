package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiumx/go-ecommerce-backend-api/internal/services"
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
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetUserByIdService(id),
	})
}
