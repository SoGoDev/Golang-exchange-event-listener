package pool

import (
	"Golang-exchange-event-listener/internal/bitmex"
	"Golang-exchange-event-listener/internal/client"
	"encoding/json"
	"fmt"
	"time"
)

type Pool struct {
	Register   chan *client.Client
	Unregister chan *client.Client
	Clients    map[*client.Client][]string
	QuitCH     chan struct{}
	Broadcast  chan []byte
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *client.Client),
		Unregister: make(chan *client.Client),
		Clients:    make(map[*client.Client][]string),
		Broadcast:  make(chan []byte),
		QuitCH:     make(chan struct{}),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case c := <-p.Register:
			p.Clients[c] = c.TrackedQuotes
			break
		case c := <-p.Unregister:
			delete(p.Clients, c)
			break
		case um := <-p.Broadcast:
			var m bitmex.Message
			_ = json.Unmarshal(um, &m)

			if m.Action != "update" {
				continue
			}

			for k, v := range p.Clients {
				for i := 0; i < len(v); i++ {
					if m.Data[0].Symbol == v[i] {
						p.SendQuotesMessage(k, m.Data[0])
					}
				}
			}

			break
		case <-p.QuitCH:
			return
		}
	}
}

func (p *Pool) Quit() {
	p.QuitCH <- struct{}{}
}

func (Pool) SendQuotesMessage(c *client.Client, data bitmex.Data) {
	qt := client.QuotesMessage{
		Timestamp: time.Now().Unix(),
		Symbol:    data.Symbol,
		Price:     data.LastPrice,
	}

	marshaled, err := json.Marshal(qt)

	if err != nil {
		fmt.Println("Failed to send quotes message")
		return
	}

	c.Write(marshaled)
}

func (p Pool) ConnectNewClient(c *client.Client) {
	p.Register <- c
}

func (p Pool) DisconnectClient(c *client.Client) {
	p.Unregister <- c
}

func (p Pool) BroadcastMessage(b []byte) {
	p.Broadcast <- b
}
