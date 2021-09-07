package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

func SendResponse(c echo.Context, data, meta, err interface{}, statusCode int) error {
	response := BaseResponse{}
	response.Data = data
	response.Meta = meta
	if err != nil {
		response.Data = err.(error).Error()
		statusCode = http.StatusInternalServerError
	}
	return c.JSON(statusCode, response)
}

func SendSucessResponse(c echo.Context, data, meta interface{}, statusCode int) error {
	response := BaseResponse{}
	response.Data = data
	response.Meta = meta
	return c.JSON(statusCode, response)
}

func SendBadResponse(c echo.Context, err interface{}, statusCode int) error {
	response := BaseResponse{}
	response.Data = err.(error).Error()
	return c.JSON(statusCode, response)
}
