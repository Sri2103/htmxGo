package main

import (
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	session "github.com/spazzymoto/echo-scs-session"
	"github.com/sri2103/htmx_go/internal/chat"
	"github.com/sri2103/htmx_go/internal/config"
	"github.com/sri2103/htmx_go/internal/pkg/database"
	"github.com/sri2103/htmx_go/internal/pkg/router"
	handler "github.com/sri2103/htmx_go/internal/todo/handlers"
	"github.com/sri2103/htmx_go/internal/todo/repository"
	userModel "github.com/sri2103/htmx_go/internal/userAuth/model"
	"github.com/sri2103/htmx_go/internal/web"
)

var sessionManager *scs.SessionManager

func main() {
	// initialize db
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.LoadConfig()
	db, err := database.ConnectSQL(cfg)
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	defer db.Conn.Close()

	// initiate a server

	server := echo.New()

	sessionManager = userModel.LoadSession()
	cfg.Server.SessionManager = sessionManager

	server.Use(session.LoadAndSave(sessionManager))

	// set Router

	Router := router.New(server)

	// start repository and handler

	repo := repository.NewRepo(db)
	todoHandler := handler.New(repo, cfg)

	// start web handlers
	webHandler := web.NewWebHandler(db, cfg)

	chatHandler := chat.NewChatServer()

	// start handlers
	Router.Run([]router.Route{
		todoHandler.AssignApiRoutes,
		webHandler.AssignWebRoutes,
		chatHandler.AssignChatRoutes,
	})

	// start server here

	log.Fatal(server.Start(cfg.Server.Addr))
}
