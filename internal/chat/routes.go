package chat

import "github.com/labstack/echo/v4"

func (ch *chatHandler) AssignChatRoutes(e *echo.Echo){
	//assign a sub Router
	Router := e.Group("/chat")
	//add the routes to the main router
	// assign websocket route to chat handler Hello
	Router.GET("/hello",ch.hello)
	// Router.("/hello",ch.hello)
}