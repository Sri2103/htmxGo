package web

import (
	"github.com/sri2103/htmx_go/internal/todo/repository"
	"github.com/sri2103/htmx_go/internal/todo/service"
	jokeService"github.com/sri2103/htmx_go/internal/jokes/service"
)

type webHandler struct {
	service service.IService
	jokeService jokeService.IService
}

func NewWebHandler(repo repository.IRepository) *webHandler {
	return &webHandler{
		service: service.New(repo),
		jokeService: jokeService.New(),
	}
}
