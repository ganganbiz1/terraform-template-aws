package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IfHealthcheckHandler interface {
	Healthcheck(c echo.Context) error
}

type HealthcheckHandler struct {
}

func NewHealthcheckHandler() IfHealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) Healthcheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
