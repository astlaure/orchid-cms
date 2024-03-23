package core

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func AuthGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)

		log.Warn(err)
		// if err != nil {
		// 	return nil
		// }

		authenticated := sess.Values["authenticated"]

		if authenticated != true {
			return c.NoContent(http.StatusUnauthorized)
		}

		return next(c)
	}
}
