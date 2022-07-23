package main

import (
	"fmt"
	"testing"
)

func TestCreateSessionToken(t *testing.T) {

	got := CreateSessionToken()
	want := "10"

	fmt.Printf(got)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
