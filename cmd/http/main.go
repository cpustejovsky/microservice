package main

import (
	"flag"
	"net/http"

	"github.com/cpustejovsky/microservice/internal/logger"
)

type Config struct {
	Addr string
}

type application struct{}

func main() {
	// Flag and Config Setup
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":5000", "HTTP network address")
	flag.Parse()

	//TODO: DB Setup?
	logger.Info.Println("Successfully connected to database!")

	// Application and Server Initialization
	app := &application{}

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: app.routes(),
	}
	logger.Info.Printf("Starting server on %s", cfg.Addr)

	// Server Start
	err := srv.ListenAndServe()
	logger.Error.Fatal(err)
}
