package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/cpustejovsky/microservice"
	"github.com/cpustejovsky/microservice/internal/logger"
)

//FormData represents the JSON data passed to CheckIPAddress handler
type FormData struct {
	IP        string   `json:"ip"`
	Whitelist []string `json:"whitelist"`
}

//Response represents the JSON data passed from the CheckIPAddress handler
type Response struct {
	WhiteListed bool   `json:"whitelisted"`
	Error       string `json:"error"`
}

func (app *application) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func (app *application) CheckIPAddress(w http.ResponseWriter, r *http.Request) {
	var data FormData
	var msg bytes.Buffer

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		logger.Error.Println(err)
	}
	if len(data.IP) < 1 {
		msg.WriteString("Did not provide correct IP address. ")
	}
	if len(data.Whitelist) < 1 {
		msg.WriteString("Did not provide a list of whitelisted countries. ")
	}

	ok, err := microservice.CheckIPAddress(data.IP, data.Whitelist)
	if err != nil {
		logger.Error.Println(err)
		msg.WriteString(err.Error())
	}
	res := Response{
		WhiteListed: ok,
		Error:       msg.String(),
	}
	bs, err := json.Marshal(res)
	if err != nil {
		logger.Error.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
	w.Write(bs)
}
