//go:build wireinject

//go:generate wire gen $GOFILE

package wire

import (
	"github.com/ganganbiz1/go-echo-gorm-template/app/config"
	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/service"
	"github.com/ganganbiz1/go-echo-gorm-template/app/handler"
	"github.com/ganganbiz1/go-echo-gorm-template/app/handler/validate"
	"github.com/ganganbiz1/go-echo-gorm-template/app/infra/repository"
	"github.com/ganganbiz1/go-echo-gorm-template/app/middleware"
	"github.com/ganganbiz1/go-echo-gorm-template/app/router"
	"github.com/ganganbiz1/go-echo-gorm-template/app/server"
	"github.com/ganganbiz1/go-echo-gorm-template/app/usecase"
	"github.com/google/wire"
)

type DIManager struct {
	Server *server.Server
	Router *router.Router
}

func DI() (*DIManager, func(), error) {
	wire.Build(
		wire.Struct(new(DIManager), "*"),

		server.NewServer,

		router.NewRouter,

		config.NewServerConfig,
		config.NewPrimaryDBConfig,
		config.NewReplicaDBConfig,

		middleware.NewCorsMiddleware,
		middleware.NewUserAuthMiddleware,

		handler.NewErrorHandler,
		handler.NewHealthcheckHandler,
		handler.NewUserHandler,

		validate.NewCustomValidator,

		usecase.NewUserUsecase,

		service.NewUserService,

		repository.NewBaseRepository,
		repository.NewPrimaryDBConnect,
		repository.NewReplicaDBConnect,
		repository.NewUserRepository,
	)
	return &DIManager{}, nil, nil
}
