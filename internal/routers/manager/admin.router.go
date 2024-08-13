package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {

	adminPublicRouter := Router.Group("/admin")
	{
		adminPublicRouter.POST("/login")
	}

	adminPrivateRouter := Router.Group("/admin")
	// userPrivateRouter.Use(Limiter())
	// userPrivateRouter.Use(Authentication())
	// userPrivateRouter.Use(Permission())
	{
		adminPrivateRouter.GET("/info")
	}
}
