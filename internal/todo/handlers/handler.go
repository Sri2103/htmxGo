package handler

import (
	"github.com/alexedwards/scs/v2"
	"github.com/sri2103/htmx_go/internal/config"
	"github.com/sri2103/htmx_go/internal/todo/repository"
	"github.com/sri2103/htmx_go/internal/todo/service"
)

type Handler struct {
	Service service.IService
}

var session *scs.SessionManager

func New(r repository.IRepository, cfg *config.AppConfig) *Handler {
	session = cfg.Server.SessionManager
	return &Handler{
		Service: service.New(r),
	}
}
