package main

import (
	"net/http"
	"os"

	"github.com/astlaure/orchid-cms/internal/auth"
	"github.com/astlaure/orchid-cms/internal/core"
	"github.com/astlaure/orchid-cms/internal/users"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var app = echo.New()

	app.Debug = core.GetDebugConfig()
	app.Renderer = core.GetRenderer()
	app.Validator = core.GetValidator()
	app.Static("/", "public")

	app.Pre(middleware.AddTrailingSlash())
	app.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))))

	core.ConnectToDatabase()

	auth.RegisterGroup("/api/auth", app)
	users.RegisterGroup("/api/users", app)

	app.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", &echo.Map{
			"Title": "Hello World",
		})
	})

	app.Logger.Fatal(app.Start("127.0.0.1:3000"))
}
