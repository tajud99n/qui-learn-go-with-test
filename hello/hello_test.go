package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string)  {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("dev", "")
		want := "Hello, dev"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T)  {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T)  {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Henry", "French")
		want := "Bonjour, Henry"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in yoruba", func(t *testing.T)  {
		got := Hello("Tajudeen", "Yoruba")
		want := "E len, Tajudeen"
		assertCorrectMessage(t, got, want)
	})
}