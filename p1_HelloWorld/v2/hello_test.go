package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Zyy")
	want := "Hello, Zyy"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}