package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tim")

	got := buffer.String()
	want := "Hello, Tim"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
