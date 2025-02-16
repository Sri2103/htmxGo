package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *chatHandler) hello(c echo.Context) error {
	//    check if the connection is websocket based
	if !c.IsWebSocket() {
		return echo.NewHTTPError(http.StatusBadRequest, "expected websocket")
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		msg := new(Message)

		err := ws.ReadJSON(&msg)
		if err != nil {
			return err
		}

		res := &Message{
			Name: "Server",
			Text: msg.Text + " [from " + msg.Name + "]",
		}
		err = ws.WriteJSON(res)
		if err != nil {
			return err
		}
	}
}

// multi use chat from websocket and htmx?
