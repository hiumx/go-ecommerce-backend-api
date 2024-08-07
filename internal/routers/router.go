package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/hiumx/go-ecommerce-backend-api/internal/controllers"
	"github.com/hiumx/go-ecommerce-backend-api/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	//middleware
	r.Use(middlewares.AuthMiddleware(), AA(), BB(), CC)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:uid", c.NewUserController().GetUserById)
	}

	return r
}
