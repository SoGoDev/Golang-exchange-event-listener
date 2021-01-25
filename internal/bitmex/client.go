package bitmex

import (
	"context"
	"github.com/gorilla/websocket"
	"net/url"
)

type Client struct {
	Key    string
	Secret string
	URL    string
}

func Listener() {
	d := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	u := url.URL{
		Host: "wss://testnet.bitmex.com",
		Path: "/realtime",
	}

	c := context.Background()
	d.DialContext(c, "", nil)
}
