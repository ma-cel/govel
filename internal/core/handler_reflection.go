package core

import (
	"govel/internal/core/interfaces"
	"reflect"

	"github.com/labstack/echo/v4"
)

type HandlerFunc func(c echo.Context) error

func Handle(controller interfaces.Controller, method string) echo.HandlerFunc {
	return func(c echo.Context) error {
		controllerType := reflect.TypeOf(controller)
		controllerValue := reflect.New(controllerType).Elem()

		methodValue := controllerValue.MethodByName(method)

		if !methodValue.IsValid() {
			return echo.NewHTTPError(500, "Controller method not found")
		}

		result := methodValue.Call([]reflect.Value{reflect.ValueOf(c)})

		if len(result) > 0 && !result[0].IsNil() {
			return result[0].Interface().(error)
		}

		return nil
	}
}
