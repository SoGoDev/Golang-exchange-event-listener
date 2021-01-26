package bitmex

import (
	"context"
	"github.com/gorilla/websocket"
	"net/url"
)

type Broadcaster interface {
	BroadcastMessage(b []byte)
}

type Client struct {
	Key    string
	Secret string
	URL    string
}

func Listen(b Broadcaster) {
	d := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	u := url.URL{
		Host:   "testnet.bitmex.com",
		Path:   "/realtime",
		Scheme: "wss",
	}

	q := u.Query()
	q.Add("subscribe", "instrument")

	u.RawQuery = q.Encode()

	c := context.Background()

	conn, _, _ := d.DialContext(c, u.String(), nil)

	for {
		_, payload, _ := conn.ReadMessage()
		b.BroadcastMessage(payload)
	}
}
