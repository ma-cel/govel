package main

import (
	"govel/http"
	"govel/http/middleware"
	"govel/internal/core"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env, did you rename .env.example?")
	}

	e := &core.GovelEcho{Echo: echo.New()}
	http.RegisterRoutes(e)
	middleware.RegisterGlobalMiddleware(e)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
