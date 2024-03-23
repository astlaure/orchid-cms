package main

import (
	"os"

	"github.com/astlaure/orchid-cms/internal/admin"
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

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))))

	core.ConnectToDatabase()

	auth.RegisterGroup("", app)
	users.RegisterGroup("/api/users", app)
	admin.RegisterGroup("", app)

	app.Logger.Fatal(app.Start("127.0.0.1:3000"))
}
