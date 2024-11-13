package main

import (
	"govel/http"
	"govel/internal/core"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	e := &core.GovelEcho{Echo: echo.New()}
	http.RegisterRoutes(e)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
