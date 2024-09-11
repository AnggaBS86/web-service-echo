package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct{}

const (
	SUCCESS = "success"
	FAILED  = "failed"
	ERROR   = "error"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func (c *Controller) Success(e echo.Context, data interface{}) error {
	return e.JSON(http.StatusOK, Response{
		Status: SUCCESS,
		Data:   data,
	})
}

func (c *Controller) Fail(e echo.Context, data interface{}) error {
	return e.JSON(http.StatusUnprocessableEntity, Response{
		Status: FAILED,
		Data:   data,
	})
}

func (c *Controller) Error(e echo.Context, data interface{}) error {
	return e.JSON(http.StatusInternalServerError, Response{
		Status: FAILED,
		Data:   data,
	})
}
