package todo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// test AddTodo RouteHandler

func Test_handler_AddTodo(t *testing.T) {
	td := `{"id":1,"title":"todo-1","desc":"New-Todo","status":"pending"}`

	t.Run("Test Handler Add Todo at no data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/todo", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)
		c.SetPath("/api/todo")
		err := RouteHandler.AddTodo(c)
		assert.Error(t, err)
	})

	t.Run("Test Handler Add Todo with data", func(t *testing.T) {
		jsonStr := []byte(td)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)
		c.SetPath("/api/todos")
		err := RouteHandler.AddTodo(c)
		assert.NoError(t, err)
	})
}
