package main

import (
	"net/http"
	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/api", app.HelloWorld)
	mux.Post("/api/:ip", app.CheckIPAddress)
}
