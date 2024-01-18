package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sri2103/htmx_go/internal/todo/model"
)

func (w *webHandler) Home(c echo.Context) error {

	data := make(map[string]interface{})

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

	err := c.Render(http.StatusOK, "home.page", data)

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
	if err != nil {
		return echo.NewHTTPError(500, "error getting todos: "+err.Error())
	}
	return c.Render(200, "todoList", todos)
}


func (w *webHandler) ShowEdit(c echo.Context) error {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	todo,err := w.service.GetTodoById(c.Request().Context(),i)
	fmt.Println(*todo)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(404,"Not Found!")
	}
	return c.Render(http.StatusOK, "todoEditCard", todo)
}

func (w *webHandler) GetSingleTodo(c echo.Context) error {
	id := c.Param("id")
	tid,_:=strconv.ParseInt(id, 10,64)
	todo,err:=w.service.GetTodoById(c.Request().Context(), int(tid))
	if err!=nil{
		return echo.NewHTTPError(404, "No Todo found with ID: " + err.Error())
	}
	return c.Render(200, "todoCard", todo)
}

func (w *webHandler) UpdateTodo(c echo.Context) error {
	t := new(model.Todo)
	title := c.FormValue("title")
	desc := c.FormValue("desc")
	t.Title = title
	t.Description = desc
	id ,_ := strconv.Atoi(c.Param("id"))
	t.ID = id

	err := w.service.UpdateTodo(id, t)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to update the data: "+err.Error())
		}
	
	return c.Render(200,"todoCard",t)
}