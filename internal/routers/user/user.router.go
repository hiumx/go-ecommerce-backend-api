package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hiumx/go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandler()

	userPublicRouter := Router.Group("/user", userController.Register)
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
