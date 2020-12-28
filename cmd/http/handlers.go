package main

import (
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

//TODO: best practice for microservice responses
//Response represents the JSON data passed from the CheckIPAddress handler
type Response struct {
	WhiteListed bool     `json:"whitelisted"`
	Message     []string `json:"message"`
	Error       string   `json:"error"`
}

func (app *application) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

//TODO: Should I change name of microservices.CheckIPAddress to clarify which is which?
func (app *application) CheckIPAddress(w http.ResponseWriter, r *http.Request) {
	var data FormData
	var msg []string
	
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		logger.Error.Println(err)
		msg = append(msg, "failed to decode body")
	}
	//TODO: is there a better, cleaner way to handle all this potential information I want to add to response?
	if len(data.IP) < 1 {
		msg = append(msg, "did not provide correct IP address")
	}
	if len(data.Whitelist) < 1 {
		msg = append(msg, "did not provide a list of whitelisted countries")
	}
	
	ok, err := microservice.CheckIPAddress(data.IP, data.Whitelist)
	if err != nil {
		logger.Error.Println(err)
		msg = append(msg, "IP Address does not match whitelisted countries")
	}
	if ok {
		msg = append(msg, "IP Address matches whitelisted countries")
	}
	res := Response{
		WhiteListed: ok,
		Message:     msg,
		Error:       err.Error(),
	}
	bs, err := json.Marshal(res)
	if err != nil {
		logger.Error.Println(err)
		//TODO: What error should be returned when json.Marshall fails?
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
	w.Write(bs)
}
