package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping/:name", Pong)
		v1.GET("/pong", Pong)
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
