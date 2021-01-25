package main

import (
	"Golang-exchange-event-listener/api/server"
	"log"
)

func main() {
	s := server.SetupServer()

	err := s.Run("localhost:12312")
	if err != nil {
		log.Fatal(err)
	}
}
