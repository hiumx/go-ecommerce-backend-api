//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/hiumx/go-ecommerce-backend-api/internal/controllers"
	"github.com/hiumx/go-ecommerce-backend-api/internal/repo"
	"github.com/hiumx/go-ecommerce-backend-api/internal/services"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		services.NewUserService,
		controllers.NewUserController,
	)

	return new(controllers.UserController), nil
}
