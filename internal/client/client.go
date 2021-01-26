package client

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn          *websocket.Conn
	TrackedQuotes []string
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{Conn: conn}
}

func (c Client) Write(data []byte) {
	err := c.Conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		fmt.Println(err)
	}
}
