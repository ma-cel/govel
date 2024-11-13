package middleware

import (
	"github.com/labstack/echo/v4/middleware"
	"goopher/internal/core"
)

func RegisterGlobalMiddleware(e *core.GovelEcho) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
