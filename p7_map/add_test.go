package main

import "testing"

func TestAdd(t *testing.T) {
	d := Dictionary{}
	d.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := d.Search("test")
	
	if err != nil {
		t.Fatal("should find adder word:", err)
	}
	
	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}