package ip

import (
	"log"
	"net"

	"github.com/cpustejovsky/microservice/internal/logger"
	"github.com/oschwald/geoip2-golang"
)

func FindCountryByIP(ipaddr string) string{
	db, err := geoip2.Open("/home/cpustejovsky/development/microservice/internal/ip/GeoLite2-Country.mmdb")
	if err != nil {
		logger.Error.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ipaddr)
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	return record.Country.Names["en"]
}
