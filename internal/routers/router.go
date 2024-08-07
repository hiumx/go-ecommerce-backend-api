package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/hiumx/go-ecommerce-backend-api/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:uid", c.NewUserController().GetUserById)
	}

	return r
}
