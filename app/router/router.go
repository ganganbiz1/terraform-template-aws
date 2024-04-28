package router

import (
	"github.com/ganganbiz1/go-echo-gorm-template/app/handler"
	"github.com/ganganbiz1/go-echo-gorm-template/app/middleware"
	"github.com/labstack/echo/v4"
)

type Router struct {
	corsMiddleware     *middleware.CorsMiddleware
	userAuthMiddleware *middleware.UserAuthMiddleware
	healthcheckHandler handler.IfHealthcheckHandler
	userHandler        handler.IfUserHandler
}

func NewRouter(
	corsMiddleware *middleware.CorsMiddleware,
	userAuthMiddleware *middleware.UserAuthMiddleware,
	healthcheckHandler handler.IfHealthcheckHandler,
	userHandler handler.IfUserHandler,

) *Router {
	return &Router{
		corsMiddleware,
		userAuthMiddleware,
		healthcheckHandler,
		userHandler,
	}
}

func (r *Router) Apply(e *echo.Echo) {
	g := e.Group("")
	group(g, "", []echo.MiddlewareFunc{}, func(g *echo.Group) {
		g.GET("/healthcheck", r.healthcheckHandler.Healthcheck)
		group(g.Group("/workers", r.userAuthMiddleware.Log), "", []echo.MiddlewareFunc{}, func(g *echo.Group) {
			g.POST("", r.userHandler.Signup)
			g.GET("/:id", r.userHandler.Search)
		})
	})

}

func group(g *echo.Group, path string, ms []echo.MiddlewareFunc, f func(*echo.Group)) {
	f(g.Group(path, ms...))
}
