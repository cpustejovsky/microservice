package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/api", http.HandlerFunc(app.HelloWorld))
	mux.Post("/api/checkip", http.HandlerFunc(app.CheckIPAddress))
	return mux
}
