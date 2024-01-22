package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sri2103/htmx_go/internal/todo/model"
	"github.com/sri2103/htmx_go/internal/web/helpers"
	"github.com/sri2103/htmx_go/internal/web/models"
)

func (w *webHandler) Home(c echo.Context) error {

	data := make(map[string]any)

	// data["title"] = "Hello"
	// data["description"] = "Lets describe it"
	// data["PageTitle"] = "Todo"
	data["todos"] = []struct {
		Title       string
		Description string
	}{
		{
			Title:       "Buy groceries",
			Description: "Milk, bread, cheese and so on...",
		},
		{
			Title:       "Clean the house",
			Description: "Kitchen, living room...",
		},
		{
			Title:       "Go to gym",
			Description: "Go to gym daily",
		},
	}

	data["PageTitle"] = "todos"

	err := c.Render(http.StatusOK, "pages/home", &models.TemplateData{
		Data: data,
	})

	if err != nil {
		fmt.Println(err.Error())
		return err

	}

	return nil
}

func (w *webHandler) AddTodo(c echo.Context) error {
	t := &model.Todo{}
	title := c.FormValue("title")
	desc := c.FormValue("desc")
	t.Title = title
	t.Description = desc

	id, err := w.service.CreateTodo(c.Request().Context(), t)
	if err != nil {
		return echo.NewHTTPError(500, "error adding todo: "+err.Error())
	}
	t.ID = id
	return c.Render(http.StatusAccepted, "todoCard", t)
}

func (w *webHandler) GetTodos(c echo.Context) error {
	todos, err := w.service.ReadTodos()

	isLogin := helpers.IsAuthenticated(c)
	fmt.Println(isLogin, "LoggedIn users")
	if err != nil {
		return echo.NewHTTPError(500, "error getting todos: "+err.Error())
	}
	data := make(map[string]interface{})
	data["todos"] = todos
	return c.Render(200, "todoList", &models.TemplateData{
		Data: data,
	})
}

func (w *webHandler) ShowEdit(c echo.Context) error {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	todo, err := w.service.GetTodoById(c.Request().Context(), i)
	fmt.Println(*todo)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(404, "Not Found!")
	}
	data := make(map[string]interface{})
	data["todo"] = todo
	return c.Render(http.StatusOK, "todoEditCard", &models.TemplateData{
		Data: data,
	})
}


func (w *webHandler) GetSingleTodo(c echo.Context) error {
	id := c.Param("id")
	tid, _ := strconv.ParseInt(id, 10, 64)
	todo, err := w.service.GetTodoById(c.Request().Context(), int(tid))
	if err != nil {
		return echo.NewHTTPError(404, "No Todo found with ID: "+err.Error())
	}
	return c.Render(200, "todoCard", todo)
}

func (w *webHandler) UpdateTodo(c echo.Context) error {
	t := new(model.Todo)
	title := c.FormValue("title")
	desc := c.FormValue("desc")
	t.Title = title
	t.Description = desc
	id, _ := strconv.Atoi(c.Param("id"))
	t.ID = id

	err := w.service.UpdateTodo(id, t)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to update the data: "+err.Error())
	}

	return c.Render(200, "todoCard", t)
}

func (w *webHandler) DummyServerPage(c echo.Context) error {
	// joke, err := w.jokeService.GetRandomJoke()
	// if err != nil {
	// 	return err
	// }
	var d = make(map[string]interface{})
	d["PageTitle"] = "served"

	e := c.Render(200, "pages/server", &models.TemplateData{
		Data: d,
	})
	return e
}

func (w *webHandler) DummyServerHandler(c echo.Context) error {
	var d = make(map[string]interface{})
	d["PageTitle"] = "dummy"
	e := c.Render(200, "pages/dummy", &models.TemplateData{
		Data: d,
	})
	return e
}

func (w *webHandler) DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := w.service.DeleteTodo(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(400, "Failed to delete todo :"+err.Error())
	}
	return c.NoContent(http.StatusOK)
}
