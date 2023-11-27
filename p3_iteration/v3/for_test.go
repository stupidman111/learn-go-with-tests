package main

import (
	"fmt"
	"testing"
)


func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}	

//
func ExampleRepeat() {
	repeated := Repeat("a", 4)
	fmt.Println(repeated)
	// Output: aaaa
}