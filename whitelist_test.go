package microservice_test

import (
	"errors"
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
		got, err := microservice.CheckIPAddress(testIPAddr, testWhiteList)
		if err != nil {
			t.Error(err)
		}
		test.Compare(t, got, want)
	})
	t.Run("Returns false when IP address is not in whitelist", func(t *testing.T) {
		testWhiteList := []string{
			"United States",
			"Mexico",
		}

		want := false
		got, err := microservice.CheckIPAddress(testIPAddr, testWhiteList)
		if err != nil {
			t.Error(err)
		}
		test.Compare(t, got, want)
	})
	t.Run("Returns false when there is an error present", func(t *testing.T) {
		testWhiteList := []string{
			"United Kingdom",
			"United States",
			"Mexico",
		}

		want := false
		got, _ := microservice.CheckIPAddress("1", testWhiteList)
		test.Compare(t, got, want)
	})
	t.Run("Returns false when there is an error present", func(t *testing.T) {
		testWhiteList := []string{
			"United Kingdom",
			"United States",
			"Mexico",
		}

		_, err := microservice.CheckIPAddress("1", testWhiteList)
		if err != nil {
			test.Compare(t, err, errors.New("1 is an incorrectly formatted IP Address"))
		}
	})

}
