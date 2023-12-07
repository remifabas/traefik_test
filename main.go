package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)

	e.GET("/status", status)
	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func status(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
