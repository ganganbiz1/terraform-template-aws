package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type CorsMiddleware struct {
}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (m *CorsMiddleware) Setting() echo.MiddlewareFunc {
	return m.setting()
}

func (m *CorsMiddleware) setting() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", os.Getenv("WOKER_HOST"))
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			if c.Request().Method == http.MethodOptions {
				return c.NoContent(http.StatusOK)
			}
			return next(c)
		}
	}
}
