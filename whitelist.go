package microservice

import (
	"github.com/cpustejovsky/microservice/internal/ip"
)

//CheckIPAddress takes an IP address as a string and a whitelist as a slice of strings. it returns a bool based on whether the string exists
func CheckIPAddress(ipaddr string, whitelist []string) (bool, error) {
	country, err := ip.FindCountryByIP(ipaddr)
	if err != nil {
		return false, err
	}
	for _, v := range whitelist {
		if v == country {
			return true, nil
		}
	}
	return false, err
}
