package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pimenvibritania/rest-api/controllers"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, This echo")
	})

	e.GET("/posts", controllers.FetchAllPost)

	return e
}
