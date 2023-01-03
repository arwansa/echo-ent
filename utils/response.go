package utils

import "github.com/labstack/echo"

type response struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func ReturnResponse(c echo.Context, code int, err error, data interface{}) error {
	response := response{
		Status: true,
		Code:   code,
		Data:   data,
	}
	if err != nil {
		response.Status = false
		response.Error = err.Error()
	}
	return c.JSON(code, response)
}
