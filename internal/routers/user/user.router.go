package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userPublicRouter := Router.Group("/user")
	{
		userPublicRouter.POST("/register")
		userPublicRouter.POST("/otp")
	}

	userPrivateRouter := Router.Group("/user")
	// userPrivateRouter.Use(Limiter())
	// userPrivateRouter.Use(Authentication())
	// userPrivateRouter.Use(Permission())
	{
		userPrivateRouter.POST("/get_info")
	}
}
