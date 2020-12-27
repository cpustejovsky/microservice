package microservice

import (
	"github.com/cpustejovsky/microservice/internal/ip"
	"github.com/cpustejovsky/microservice/internal/logger"
)

//CheckIPAddress takes an IP address as a string and a whitelist as a slice of strings. it returns a bool based on whether the string exists
func CheckIPAddress(ipaddr string, whitelist []string) bool {
	country, err := ip.FindCountryByIP(ipaddr)
	if err != nil {
		logger.Error.Println(err)
		return false
	}
	for _, v := range whitelist {
		if v == country {
			return true
		}
	}
	return false
}
