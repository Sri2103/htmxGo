package handler

import "github.com/labstack/echo/v4"

func (h *Handler) AssignApiRoutes(e *echo.Echo) {
	router := e.Group("/api")
	router.POST("/todos", h.AddTodo)
	router.GET("/todos", h.GetAllTodo)
	router.DELETE("/todos/:id", h.DeleteTodo)
	router.GET("/todos/:id", h.GetSingleTodo)
}
