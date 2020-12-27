package main

import (
	"github.com/cpustejovsky/microservice/internal/ip"
	"github.com/cpustejovsky/microservice/internal/logger"
)

func main() {
	logger.Info.Println("Hello, World!")
	country := ip.FindCountryByIP("81.2.69.142")
	logger.Info.Println("COUNTRY:\t", country)
}
