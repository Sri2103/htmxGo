package router

import "github.com/labstack/echo/v4"

type Router struct {
	server *echo.Echo
}

type Route func(e *echo.Echo)


func New(e *echo.Echo) *Router {
	return &Router{
		server: e,
	}
}

func (e *Router) setRouter(routes []Route) {
	for _, r := range routes {
		r(e.server)
	}
}

func (e *Router) Run(r []Route) {
	e.setRouter(r)
}

