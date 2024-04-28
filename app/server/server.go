package server

import (
	"fmt"

	"github.com/ganganbiz1/go-echo-gorm-template/app/config"
	"github.com/ganganbiz1/go-echo-gorm-template/app/handler"
	"github.com/ganganbiz1/go-echo-gorm-template/app/handler/validate"
	"github.com/ganganbiz1/go-echo-gorm-template/app/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo         *echo.Echo
	serverConfig *config.ServerConfig
}

func NewServer(
	serverConfig *config.ServerConfig,
	corsMiddleware *middleware.CorsMiddleware,
	errHandler *handler.ErrorHandler,
	validator *validate.CustomValidator,
) *Server {
	e := echo.New()

	e.HTTPErrorHandler = errHandler.Handler
	e.Validator = validator

	e.Use(corsMiddleware.Setting())
	e.Use(echoMiddleware.Recover())
	return &Server{
		Echo:         e,
		serverConfig: serverConfig,
	}
}

func (s *Server) Start() error {
	return s.Echo.Start(fmt.Sprintf(":%d", s.serverConfig.Port))
}
