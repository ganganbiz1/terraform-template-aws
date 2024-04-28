package presenter

import (
	"net/http"

	"github.com/ganganbiz1/go-echo-gorm-template/app/domain/entity"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID            int    `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	CognitoUserID string `json:"cognitoUserID"`
}

func BuildUserSearchResponse(c echo.Context, e *entity.User) error {
	return BuildSuccessResponse(c, http.StatusOK, &User{
		ID:            e.ID,
		Email:         e.Email,
		Name:          e.Name,
		CognitoUserID: e.CognitoUserID,
	})
}
