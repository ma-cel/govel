package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ExampleController struct{}

func NewExampleController() *ExampleController {
	return &ExampleController{}
}

func (c *ExampleController) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Hello from ExampleController"})
}
