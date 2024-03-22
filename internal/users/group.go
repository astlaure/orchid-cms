package users

import (
	"net/http"

	"github.com/astlaure/orchid-cms/internal/core"
	"github.com/labstack/echo/v4"
)

func RegisterGroup(prefix string, app *echo.Echo) {
	var group = app.Group(prefix)

	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	core.DB.AutoMigrate(&User{})
}
