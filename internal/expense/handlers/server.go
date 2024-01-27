package handlers

import service "github.com/sri2103/htmx_go/internal/expense/services"

type Handler struct {
	Service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{
		Service: service,
		}
	}