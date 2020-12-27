package main

import (
	"encoding/json"
	"net/http"

	"github.com/cpustejovsky/microservice"
	"github.com/cpustejovsky/microservice/internal/logger"
)

func (app *application) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

//TODO: Should I change name of microservices.CheckIPAddress to clarify which is which?
func (app *application) CheckIPAddress(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type FormData struct {
		IP        string   `json:"ip"`
		Whitelist []string `json:"whitelist"`
	}
	var data FormData
	err := decoder.Decode(&data)
	if err != nil {
		logger.Error.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}
	ok := microservice.CheckIPAddress(data.IP, data.Whitelist)

	type Response struct {
		Message string `json:"message"`
		Value   bool   `json:"value"`
	}
	res := Response{
		Message: "success",
		Value:   ok,
	}
	bs, err := json.Marshal(res)
	if err != nil {
		logger.Error.Println(err)
		//TODO: What error should be returned when json.Marshall fails?
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
	w.Write(bs)
}
