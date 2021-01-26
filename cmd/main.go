package main

import (
	"Golang-exchange-event-listener/api/server"
	"Golang-exchange-event-listener/internal/bitmex"
	"Golang-exchange-event-listener/internal/pool"
	"log"
)

func main() {
	p := pool.NewPool()
	go p.Start()

	defer p.Quit()

	go bitmex.Listen(p)

	s := server.SetupServer(p)

	err := s.Run("localhost:12312")
	if err != nil {
		log.Fatal(err)
	}
}
