package http

import (
	"govel/controllers"
	"govel/internal/core"
)

func RegisterRoutes(gv *core.GovelEcho) {
	gv.GET("/", &controllers.ExampleController{}, "Index")
}
