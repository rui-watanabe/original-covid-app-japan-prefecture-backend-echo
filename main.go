package main

import (
	"fmt"
	"net/http"
	"original-covid-app-japan-prefecture-backend-echo/config"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("hello world")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":" + config.Config.Port))
}
