package ip

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/oschwald/geoip2-golang"
)

//FindCountryByIP takes an IP address in the form of a string and returns the English name of the country that corresponds to the IP address and an error
func FindCountryByIP(ipaddr string) (string, error) {
	if err := godotenv.Load(os.ExpandEnv("$GOPATH/src/microservice/.env")); err != nil {
		return "", err
	}
	geodb := os.Getenv("GEO_MMDB")
	db, err := geoip2.Open(geodb)
	if err != nil {
		return "", err
	}
	defer db.Close()
	ip := net.ParseIP(ipaddr)
	if ip == nil {
		return "", fmt.Errorf("%v is an incorrectly formatted IP Address", ipaddr)
	}
	record, err := db.Country(ip)
	if err != nil {
		return "", err
	}
	return record.Country.Names["en"], nil
}
