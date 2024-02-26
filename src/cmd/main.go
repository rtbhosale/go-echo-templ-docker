package main

import (
	template "github.com/rtbhosale/go-echo-templ-docker/template/index"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return template.Index().Render(c.Request().Context(), c.Response().Writer)
	})
	e.Logger.Fatal(e.Start(":1313"))

}
