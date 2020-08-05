package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Goutham", "")
		want := "Hello, Goutham"

		assertCorrectMessage(t, got, want)
	})

	t.Run("sat 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Goutham", "Spanish")
		want := "Hola, Goutham"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Goutham", "French")
		want := "Bonjour, Goutham"

		assertCorrectMessage(t, got, want)
	})
}
