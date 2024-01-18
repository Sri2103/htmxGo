package web

import (
	"github.com/sri2103/htmx_go/internal/todo/repository"
	"github.com/sri2103/htmx_go/internal/todo/service"
)

type webHandler struct {
	service service.IService
}

func NewWebHandler(repo repository.IRepository) *webHandler {
	return &webHandler{
		service: service.New(repo)}
}
