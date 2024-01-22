package web

import (
	"github.com/alexedwards/scs/v2"
	"github.com/sri2103/htmx_go/internal/config"
	jokeService "github.com/sri2103/htmx_go/internal/jokes/service"
	"github.com/sri2103/htmx_go/internal/pkg/database"
	"github.com/sri2103/htmx_go/internal/todo/repository"
	"github.com/sri2103/htmx_go/internal/todo/service"
	userRepo "github.com/sri2103/htmx_go/internal/userAuth/repository"
	userService "github.com/sri2103/htmx_go/internal/userAuth/service"
	"github.com/sri2103/htmx_go/internal/web/helpers"
)

type webHandler struct {
	session     *scs.SessionManager
	service     service.IService
	jokeService jokeService.IService
	userService userService.IService
}

func NewWebHandler(db *database.DB, cfg *config.AppConfig) *webHandler {
	helpers.InitiateHelpers(cfg)
	return &webHandler{
		session:     cfg.Server.SessionManager,
		service:     service.New(repository.NewRepo(db)),
		jokeService: jokeService.New(),
		userService: userService.NewUserService(userRepo.NewRepo(db)),
	}
}
