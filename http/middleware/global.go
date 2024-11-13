package middleware

import (
	"govel/internal/core"

	"github.com/labstack/echo/v4/middleware"
)

func RegisterGlobalMiddleware(e *core.GovelEcho) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
