package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hiumx/go-ecommerce-backend-api/internal/controllers"
	"github.com/hiumx/go-ecommerce-backend-api/internal/repo"
	"github.com/hiumx/go-ecommerce-backend-api/internal/services"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	// No dependency injection
	userRepo := repo.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

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
