package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// initialise echo server
	e := echo.New()
	e.HideBanner = true

	// html templating
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	// echo renderer
	e.Renderer = t

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// index route
	e.GET("/", index)

	// health check route
	e.GET("/health", handleHealthCheck)

	// start the echo server
	e.Logger.Fatal(e.Start(":8000"))
}

// implement echo.Renderer interface
type Template struct {
	templates *template.Template
}

// implement echo.Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// index handler
func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "World")
}

// health check handler
func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
