package todo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetHandler(t *testing.T) {
	t.Run("Handler getTodos", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		c := srv.NewContext(req, w)
		c.SetPath("/api/todo")
		err := RouteHandler.GetAllTodo(c)
		assert.NoError(t, err)
	})
}
