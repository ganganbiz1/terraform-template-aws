package middleware

import (
	"github.com/ganganbiz1/go-echo-gorm-template/app/logger"
	"github.com/labstack/echo/v4"
)

type UserAuthMiddleware struct {
}

func NewUserAuthMiddleware() *UserAuthMiddleware {
	return &UserAuthMiddleware{}
}

func (m *UserAuthMiddleware) Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.Info("UserAuthMiddleware exec")
		return next(c)
	}
}
