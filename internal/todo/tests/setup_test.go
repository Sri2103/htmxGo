package todo

import (
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sri2103/htmx_go/internal/pkg/router"
	handler "github.com/sri2103/htmx_go/internal/todo/handlers"
	"github.com/sri2103/htmx_go/internal/todo/repository"
)

var (
	RouteHandler *handler.Handler
	srv *echo.Echo
)

func TestMain(m *testing.M){
	// setup code here
	srv = echo.New()
	repo := repository.NewTestRepo()
	repo.CreateRecords()

	RouteHandler = handler.New(repo)

	Router := router.New(srv)

	Router.Run([]router.Route{
		RouteHandler.AssignApiRoutes,
	})
	code := m.Run()
	// teardown code here
	os.Exit(code)
}