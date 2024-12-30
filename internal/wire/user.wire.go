//go:build wireinject
// +build wireinject

// annotate the file with the wireinject build tag to enable wire to generate wire_gen.go

package wire

import (
	"github.com/google/wire"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/controller"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/repo"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return &controller.UserController{}, nil
}
