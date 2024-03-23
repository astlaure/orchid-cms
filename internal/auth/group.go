package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func RegisterGroup(prefix string, app *echo.Echo) {
	var group = app.Group(prefix)

	group.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", echo.Map{})
	})

	group.POST("/login", func(c echo.Context) error {
		var body Login
		err := c.Bind(&body)

		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		err = validateUser(body)

		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["authenticated"] = true
		sess.Save(c.Request(), c.Response())

		return c.NoContent(http.StatusOK)
	})

	group.POST("/logout", func(c echo.Context) error {
		sess, _ := session.Get("session", c)

		sess.Values["authenticated"] = false
		sess.Save(c.Request(), c.Response())

		return c.NoContent(http.StatusOK)
	})
}
