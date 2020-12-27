package main

import (
	"github.com/cpustejovsky/microservice/internal/ip"
	"github.com/cpustejovsky/microservice/internal/logger"
)

func main() {
	logger.Info.Println("Hello, World!")
	country, err := ip.FindCountryByIP("81.2.69.1423")
	if err != nil {
		logger.Error.Fatal(err)
	}
	logger.Info.Println("COUNTRY:\t", country)
}
