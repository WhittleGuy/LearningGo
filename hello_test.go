package main

import "testing"

func TestHello (t *testing.T) {
	got := Hello("Whittle")
	want := "Hello, Whittle!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}