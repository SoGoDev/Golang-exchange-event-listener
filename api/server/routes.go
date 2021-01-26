package server

import (
	"Golang-exchange-event-listener/internal/client"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Manager interface {
	ConnectNewClient(c *client.Client)
	DisconnectClient(c *client.Client)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func SetupServer(reg Manager) *gin.Engine {
	e := gin.Default()

	e.GET("/ws", func(c *gin.Context) {
		WebSocketListener(reg)(c.Writer, c.Request)
	})

	return e
}
