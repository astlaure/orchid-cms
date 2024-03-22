package users

import (
	"net/http"
	"strconv"

	"github.com/astlaure/orchid-cms/internal/core"
	"github.com/labstack/echo/v4"
)

func RegisterGroup(prefix string, app *echo.Echo) {
	var group = app.Group(prefix)

	group.GET("/", func(c echo.Context) error {
		pageNumber, err := strconv.ParseUint(c.QueryParam("page"), 10, 32)

		if err != nil {
			pageNumber = 0
		}

		users, err := retrieveUserPage(uint(pageNumber))

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, users)
	})

	group.GET("/:id", func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err})
		}

		user, err := retrieveUserById(uint(id))

		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}

		return c.JSON(http.StatusOK, user)
	})

	group.POST("/", func(c echo.Context) error {
		var createUser CreateUser
		err := c.Bind(&createUser)

		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		user, err := insertUser(createUser)

		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		return c.JSON(http.StatusOK, user)
	})

	group.PUT("/:id", func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err})
		}

		var body UpdateUser
		err = c.Bind(&body)

		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		err = modifyUser(uint(id), body)

		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		return c.NoContent(http.StatusNoContent)
	})

	group.PATCH("/:id/password", func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err})
		}

		var body UpdateUserPassword
		err = c.Bind(&body)

		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		err = modifyUserPassword(uint(id), body)

		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		return c.NoContent(http.StatusNoContent)
	})

	group.DELETE("/:id", func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err})
		}

		err = destroyUser(uint(id))

		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		return c.NoContent(http.StatusNoContent)
	})

	core.DB.AutoMigrate(&User{})
}
