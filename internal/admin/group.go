package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterGroup(prefix string, app *echo.Echo) {
	var group = app.Group(prefix)

	group.GET("", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", &echo.Map{
			"Title": "Hello World",
		})
	})

	group.GET("/test", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", &echo.Map{
			"Title": "Hello World",
		})
	})
}
