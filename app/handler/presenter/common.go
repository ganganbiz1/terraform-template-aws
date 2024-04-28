package presenter

import "github.com/labstack/echo/v4"

type CommonResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BuildSuccessResponse(c echo.Context, httpStatus int, data interface{}) error {
	return c.JSON(httpStatus, &CommonResponse{
		Message: "success",
		Data:    data,
	})
}

func BuildErrorResponse(c echo.Context, httpStatus int, message string) error {
	return c.JSON(httpStatus, &CommonResponse{
		Message: message,
		Data:    nil,
	})
}
