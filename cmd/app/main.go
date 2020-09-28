package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/sebastianflor/golang-postgresql-api/internal/server"
)

func main() {
	s, err := server.New("8000")
	if err != nil {
		log.Fatal(err)
	}

	// start the server
	go s.Start()

	// wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// attempt a graceful shutdown
	s.Close()
}
