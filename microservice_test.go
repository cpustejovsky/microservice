package microservice_test

import (
	"testing"

	"github.com/cpustejovsky/microservice"
	"github.com/cpustejovsky/microservice/internal/test"
)

func TestCheckIPAddress(t *testing.T) {
	testIPAddr := "81.2.69.142"
	t.Run("Returns true when IP address is included", func(t *testing.T) {
		testWhiteList := []string{
			"United Kingdom",
			"United States",
			"Mexico",
		}

		want := true
		got := microservice.CheckIPAddress(testIPAddr, testWhiteList)

		test.Compare(t, got, want)
	})
	t.Run("Returns true when IP address is included", func(t *testing.T) {
		testWhiteList := []string{
			"United States",
			"Mexico",
		}

		want := false
		got := microservice.CheckIPAddress(testIPAddr, testWhiteList)

		test.Compare(t, got, want)

	})
}
