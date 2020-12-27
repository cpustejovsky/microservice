package main

import (
	"encoding/json"
	"net/http"

	"github.com/cpustejovsky/microservice/internal/logger"
)

func (app *application) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func (app *application) CheckIPAddress(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println(r.Body)
	decoder := json.NewDecoder(r.Body)
	type FormData struct {
		IP string `json:"ip"`
	}
	var data FormData
	err := decoder.Decode(&data)
	if err != nil {
		logger.Error.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}
	logger.Info.Println(data)
	w.Write([]byte(data.IP))
}
