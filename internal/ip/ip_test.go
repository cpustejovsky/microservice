package ip_test

import (
	"testing"

	"github.com/cpustejovsky/microservice/internal/ip"
	"github.com/cpustejovsky/microservice/internal/test"
)

func TestFindCountryByIP(t *testing.T) {
	want := "United Kingdom"
	got, err := ip.FindCountryByIP("81.2.69.142")
	if err != nil {
		t.Error(err)
	}
	test.Compare(t, got, want)
}
