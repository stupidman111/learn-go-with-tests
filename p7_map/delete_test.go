package main

import (
	"testing"
)

func TestDelete(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	d.Delete("test")

	_, got := d.Search("test")
	want := ErrNotFound

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}