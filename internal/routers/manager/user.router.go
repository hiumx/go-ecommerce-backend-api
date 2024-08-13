package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userPrivateRouter := Router.Group("/admin/user")
	// userPrivateRouter.Use(Limiter())
	// userPrivateRouter.Use(Authentication())
	// userPrivateRouter.Use(Permission())
	{
		userPrivateRouter.POST("/active")
	}
}
