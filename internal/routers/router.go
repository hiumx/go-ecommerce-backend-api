package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	c "github.com/hiumx/go-ecommerce-backend-api/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping/:name", Pong)
		v1.GET("/users/:uid", c.NewUserController().GetUserById)
	}

	return r
}

func Pong(c *gin.Context) {
	name := c.Param("name")

	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
		"uid":     uid,
		"users":   []string{"an", "binh", "cuong"},
	})
}
