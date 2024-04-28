package handler

import (
	"net/http"
	"strconv"

	"github.com/ganganbiz1/go-echo-gorm-template/app/handler/presenter"
	"github.com/ganganbiz1/go-echo-gorm-template/app/usecase"
	"github.com/ganganbiz1/go-echo-gorm-template/app/usecase/dto/input"
	"github.com/labstack/echo/v4"
)

type IfUserHandler interface {
	Signup(c echo.Context) error
	Search(c echo.Context) error
}

type UserHandler struct {
	userUsecase usecase.IfUserUsecase
}

type UserSignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

func NewUserHandler(userUsecase usecase.IfUserUsecase) IfUserHandler {
	return &UserHandler{
		userUsecase,
	}
}

func (h *UserHandler) Signup(c echo.Context) error {
	var req UserSignupRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	dto := &input.User{
		ID:       0,
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}
	if err := h.userUsecase.Signup(c.Request().Context(), dto); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (h *UserHandler) Search(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	e, err := h.userUsecase.Search(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return presenter.BuildUserSearchResponse(c, e)
}
