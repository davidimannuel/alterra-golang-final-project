package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

func SendSucessResponse(c echo.Context, data, meta interface{}) error {
	response := BaseResponse{}
	response.Data = data
	response.Meta = meta
	return c.JSON(http.StatusOK, response)
}

func SendBadResponse(c echo.Context, err interface{}, statusCode int) error {
	response := BaseResponse{}
	response.Data = err.(error).Error()
	return c.JSON(statusCode, response)
}
