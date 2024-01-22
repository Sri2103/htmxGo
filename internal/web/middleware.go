package web

import (
	"github.com/labstack/echo/v4"
	"github.com/sri2103/htmx_go/internal/web/helpers"
)

func AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if helpers.IsAuthenticated(ctx) {
			return next(ctx)
		}
		return ctx.Redirect(302, "/web/login")
	}
}
