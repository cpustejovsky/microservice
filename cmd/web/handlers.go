package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cpustejovsky/microservice"
	"github.com/cpustejovsky/microservice/internal/logger"
)

func (app *application) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
//TODO: change name of microservices.CheckIPAddress to clarify which is which?
func (app *application) CheckIPAddress(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println(r.Body)
	decoder := json.NewDecoder(r.Body)
	type FormData struct {
		IP        string   `json:"ip"`
		Whitelist []string `json:"whitelist`
	}
	var data FormData
	err := decoder.Decode(&data)
	if err != nil {
		logger.Error.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}
	ok := microservice.VerifyIPAddress(data.IP, data.Whitelist)

	w.Write([]byte(strconv.FormatBool(ok)))
}
