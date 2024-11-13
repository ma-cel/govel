package interfaces

import "github.com/labstack/echo/v4"

type Controller interface {
	Index(ctx echo.Context) error
}

type ResourceController interface {
	Index(ctx echo.Context) error
	Show(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
