package helpers

import (
	"github.com/labstack/echo/v4"
	"github.com/sri2103/htmx_go/internal/config"
)

var app *config.AppConfig

func InitiateHelpers(a *config.AppConfig){
	app = a
}


func IsAuthenticated(c echo.Context) bool {
	return app.Server.SessionManager.Exists(c.Request().Context(),"user")
}