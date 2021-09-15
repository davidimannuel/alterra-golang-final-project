package handlers

import (
	"keep-remind-app/businesses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
}

type Pagination struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	TotalData int `json:"total_data"`
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

func PageInfo(domain businesses.Pagination) Pagination {
	return Pagination{
		Page:      domain.Page,
		PerPage:   domain.PerPage,
		TotalData: domain.TotalData,
	}
}
