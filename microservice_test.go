package microservice_test

import (
	"testing"

	"github.com/cpustejovsky/microservice"
	"github.com/cpustejovsky/microservice/internal/test"
)

func TestCheckIPAddress(t *testing.T) {
	testIPAddr := "81.2.69.142"
	testWhiteList := []string{
		"United Kingdom",
		"United States",
		"Mexico",
	}

	want := true
	got := microservice.CheckIPAddress(testIPAddr, testWhiteList)

	test.Compare(t, got, want)
}
