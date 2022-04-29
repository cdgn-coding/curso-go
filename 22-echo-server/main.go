package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// eliminar modulos no usados
	// go mod tidy
	// agrega las carpetas de los modulos en el root
	// go mod vendor
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
