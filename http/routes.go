package http

import (
	"govel/http/controllers"
	"govel/internal/core"
)

func RegisterRoutes(gv *core.GovelEcho) {
	gv.GET("/", &controllers.ExampleController{}, "Index")
}
