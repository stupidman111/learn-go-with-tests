package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func (t *testing.T, got, want string)  {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Zyy", "Spanish")
		want := "Hola, Zyy"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In English", func(t *testing.T) {
		got := Hello("Zyy", "English")
		want := "Hello, Zyy"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In English", func(t *testing.T) {
		got := Hello("Zyy", "")
		want := "Hello, Zyy"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Zyy", "French")
		want := "Bonjour, Zyy"

		assertCorrectMessage(t, got, want)
	})
}