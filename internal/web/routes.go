package web

import (
	"fmt"
	"html/template"
	"io"

	"github.com/alexedwards/scs/v2"
	sprig "github.com/go-task/slim-sprig/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sri2103/htmx_go/internal/web/models"
	"github.com/unrolled/render"
)

type RenderWrapper struct { // We need to wrap the renderer because we need a different signature for echo.
	rnd *render.Render
	s   *scs.SessionManager
}

func (r *RenderWrapper) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// data = r.AddDefaultData(data.(*models.TemplateData), c)

	return r.rnd.HTML(w, 0, name, data) // The zero status code is overwritten by echo.
}

func (r *RenderWrapper) AddDefaultData(td *models.TemplateData, c echo.Context) *models.TemplateData {
	td.Flash = r.s.PopString(c.Request().Context(), "flash")
	td.Warning = r.s.PopString(c.Request().Context(), "warning")
	td.Error = r.s.PopString(c.Request().Context(), "error")
	fmt.Println("isAuthenticated", r.s.Exists(c.Request().Context(), "user"))
	if r.s.Exists(c.Request().Context(), "user") {
		td.IsAuthenticated = true
	} else {
		td.IsAuthenticated = false
	}
	return td
}

func (w *webHandler) AssignWebRoutes(e *echo.Echo) {
	// Get all views for the current user
	// e.Renderer = &Template{
	// 	templates: template.Must(template.New("").Funcs(sprig.FuncMap()).ParseGlob("templates/**/*.html")),
	// }
	e.Renderer = &RenderWrapper{
		rnd: render.New(render.Options{
			Funcs:      []template.FuncMap{sprig.FuncMap()},
			Extensions: []string{".tmpl", ".html"},
		}),
		s: w.session,
	}
	// e.Renderer = NewRender()
	e.Static("/static", "static")
	e.Static("/dist", "dist")
	Router := e.Group("/web")
	Router.Use(middleware.Logger())
	Router.Use(middleware.Recover())
	Router.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "cookie:_csrf",
	}))
	Router.GET("/home", w.Home, AuthRequired)
	Router.GET("/getAll", w.GetTodos, AuthRequired)
	Router.GET("/getAllDone", w.GetDoneTodos, AuthRequired)

	Router.POST("/addTodo", w.AddTodo, AuthRequired)
	Router.GET("/showEdit/:id", w.ShowEdit, AuthRequired)
	Router.GET("/todo/:id", w.GetSingleTodo, AuthRequired)
	Router.PUT("/updateTodo/:id", w.UpdateTodo, AuthRequired)
	Router.GET("/served", w.DummyServerPage, AuthRequired)
	Router.GET("/dummy", w.DummyServerHandler, AuthRequired)
	Router.DELETE("/deleteTodo/:id", w.DeleteTodo, AuthRequired)
	Router.PUT("/toggleTodo/:id", w.ToggleTodoStatus, AuthRequired)

	Router.GET("/login", w.LoginPage)
	Router.GET("/register", w.Register)
	Router.POST("/login", w.PostLogin)
	Router.POST("/register", w.PostRegister)
	Router.GET("/logout", w.LogOut, AuthRequired)
}
