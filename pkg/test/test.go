package test

import (
	"reflect"
	"testing"
)

//Compare takes T and a got and a want and checks if they are the same
func Compare(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nwanted:\n%v\ngot:\n%v\n", want, got)
	}
}
