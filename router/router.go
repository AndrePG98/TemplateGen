package router

import (
	"TemplateGen/config"
	"TemplateGen/handler"
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(app *echo.Echo, db *sql.DB, config *config.Config) {
	app.Use(middleware.CORS())
	ph := handler.NewPageHandler()
	app.GET("/", ph.HomePage)
}
