package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sebastianflor/golang-postgresql-api/internal/data"
	"github.com/sebastianflor/golang-postgresql-api/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	s, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	// connection to the database
	d := data.New()
	if err := d.DB.Ping(); err != nil {
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
	data.Close()
}
