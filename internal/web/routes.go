package web

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	sprig "github.com/go-task/slim-sprig/v3"
	"github.com/unrolled/render"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

type RenderWrapper struct { // We need to wrap the renderer because we need a different signature for echo.
    rnd *render.Render
}

func (r *RenderWrapper) Render(w io.Writer, name string, data interface{},c echo.Context) error {

    return r.rnd.HTML(w, 0, name, data) // The zero status code is overwritten by echo.
}

func (w *webHandler) AssignWebRoutes(e *echo.Echo) {
	// Get all views for the current user
	// e.Renderer = &Template{
	// 	templates: template.Must(template.New("").Funcs(sprig.FuncMap()).ParseGlob("templates/**.html")),
	// }
	e.Renderer =  &RenderWrapper{render.New(render.Options{
		Funcs: []template.FuncMap{sprig.FuncMap()},
		Extensions: []string{".tmpl",".html"},
	})}
	e.Static("/static", "static")
	Router := e.Group("/web")
	Router.Use(middleware.Logger())
	Router.GET("/home", w.Home)
	Router.GET("/getAll", w.GetTodos)
	Router.POST("/addTodo", w.AddTodo)
	Router.GET("/showEdit/:id",w.ShowEdit)
	Router.GET("/todo/:id",w.GetSingleTodo)
	Router.PUT("/updateTodo/:id",w.UpdateTodo)
}
