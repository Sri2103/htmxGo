package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sri2103/htmx_go/internal/todo/model"
)

func (h *Handler) AddTodo(c echo.Context) error {
	var t model.Todo
	if err := c.Bind(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	id, err := h.Service.CreateTodo(c.Request().Context(), &t)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "server error:"+err.Error())
	}
	t.ID = id
	return c.JSON(http.StatusCreated, t)
}

// get all todos handler
func (h *Handler) GetAllTodo(c echo.Context) error {
	id := session.Get(c.Request().Context(), "userId")
	todos, err := h.Service.ReadTodos(id.(int))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch: "+err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

// delete todo Handler

func (h *Handler) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid paramter: "+err.Error())
	}

	err = h.Service.DeleteTodo(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found: "+err.Error())
	}
	return c.String(http.StatusAccepted, "Deletion successful")
}

func (h *Handler) GetSingleTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid paramter: "+err.Error())
	}
	t, err := h.Service.GetTodoById(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found: "+err.Error())
	}
	return c.JSON(200, t)
}
