package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// initialise echo server
	e := echo.New()
	e.HideBanner = true

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// hello world route
	e.GET("/", handleHelloWorld)

	// health check route
	e.GET("/health", handleHealthCheck)

	// start the echo server
	e.Logger.Fatal(e.Start(":8000"))
}

// hello world handler
func handleHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

// health check handler
func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
