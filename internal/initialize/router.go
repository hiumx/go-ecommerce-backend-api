package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hiumx/go-ecommerce-backend-api/global"
	"github.com/hiumx/go-ecommerce-backend-api/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	userRouter := routers.RouterGroupApp.User
	managerRouter := routers.RouterGroupApp.Manager

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("check_status") // tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		managerRouter.InitAdminRouter(MainGroup)
		managerRouter.InitUserRouter(MainGroup)
	}

	return r
}
