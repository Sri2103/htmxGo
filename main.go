package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sri2103/htmx_go/internal/config"
	"github.com/sri2103/htmx_go/internal/pkg/database"
	"github.com/sri2103/htmx_go/internal/pkg/router"
	handler "github.com/sri2103/htmx_go/internal/todo/handlers"
	"github.com/sri2103/htmx_go/internal/todo/repository"
	"github.com/sri2103/htmx_go/internal/web"
)

func main() {
	// initialize db
    cfg := config.LoadConfig()
	db, err := database.ConnectSQL(&cfg)

	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	defer db.Conn.Close()

	// initiate a server

	server := echo.New()

	// set Router

	Router := router.New(server)

	// start repository and handler

	repo := repository.NewRepo(db)
	todoHandler := handler.New(repo)


	// start web handlers
	webHandler := web.NewWebHandler(repo)

	// start handlers
	Router.Run([]router.Route{
		todoHandler.AssignApiRoutes,
		webHandler.AssignWebRoutes,
	})


	// start server here

	log.Fatal(server.Start(":3500"))
}