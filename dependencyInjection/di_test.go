package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "dev")

	got := buffer.String()
	want := "Hello, dev"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
