package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"original-covid-app-japan-prefecture-backend-echo/config"
	"original-covid-app-japan-prefecture-backend-echo/openapi"
)

func main() {
	fmt.Println("hello world")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Println(openapi.Prefectures{})

	e.Logger.Fatal(e.Start(":" + config.Config.Port))
}
