package server

import (
	"Golang-exchange-event-listener/internal/client"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

const (
	ActionTypeSubscribe   = "subscribe"
	ActionTypeUnsubscribe = "unsubscribe"
)

func WebSocketListener(man Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsupgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
			return
		}

		c := client.NewClient(conn)

		for {
			var m client.Message
			err := conn.ReadJSON(&m)
			if err != nil {
				break
			}

			switch m.Action {
			case ActionTypeSubscribe:
				{
					c.TrackedQuotes = m.Symbols
					man.ConnectNewClient(c)
					break
				}
			case ActionTypeUnsubscribe:
				{
					c.Conn.Close()
					man.DisconnectClient(c)
					break
				}
			default:
				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Uknown action: %s\n", m.Action)))
			}

		}
	}
}
