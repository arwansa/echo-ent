package utils

import "github.com/labstack/echo"

type successResponse struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

type errorResponse struct {
	Status bool   `json:"status"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}

func SuccessResponse(c echo.Context, status bool, code int, data interface{}) error {
	response := successResponse{
		Status: status,
		Code:   code,
		Data:   data,
	}
	return c.JSON(code, response)
}

func ErrorResponse(c echo.Context, status bool, code int, err error) error {
	response := errorResponse{
		Status: status,
		Code:   code,
		Error:  err.Error(),
	}
	return c.JSON(code, response)
}
