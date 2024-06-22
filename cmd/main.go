package main

import (
	"TemplateGen/config"
	"TemplateGen/database"
	"TemplateGen/router"
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	config := config.LoadConfig()
	db, err := database.Init(config, 5, time.Second*3)
	if err != nil {
		slog.Error(err.Error())
	} else {
		slog.Info("Connected to the dabatase")
	}

	app := echo.New()
	app.Static("/static", "views/static")

	router.Setup(app, db, config)

	slog.Info("HTTP server started", "listenAddr", config.PORT)
	app.Logger.Fatal(app.Start(":" + config.PORT))
}
