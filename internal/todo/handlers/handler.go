package handler

import (
	"github.com/sri2103/htmx_go/internal/todo/repository"
	"github.com/sri2103/htmx_go/internal/todo/service"
)

type Handler struct {
	Service service.IService
}

func New(r repository.IRepository) *Handler {
	return &Handler{
		Service: service.New(r),
	}
}
