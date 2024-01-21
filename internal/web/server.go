package web

import (
	jokeService "github.com/sri2103/htmx_go/internal/jokes/service"
	"github.com/sri2103/htmx_go/internal/pkg/database"
	"github.com/sri2103/htmx_go/internal/todo/repository"
	userRepo "github.com/sri2103/htmx_go/internal/userAuth/repository"
	"github.com/sri2103/htmx_go/internal/todo/service"
	userService "github.com/sri2103/htmx_go/internal/userAuth/service"
)

type webHandler struct {
	service service.IService
	jokeService jokeService.IService
	userService userService.IService

}

func NewWebHandler(db *database.DB) *webHandler {
	return &webHandler{
		service: service.New(repository.NewRepo(db)),
		jokeService: jokeService.New(),
		userService: userService.NewUserService(userRepo.NewRepo(db)),
	}
}
