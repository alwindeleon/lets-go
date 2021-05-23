package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Oi, JoniBoi' if language is filipino", func(t *testing.T) {
		got := Hello("JoniBoi", "filipino")
		want := "Oi, JoniBoi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Holla, JoniBoi' if language is spanish", func(t *testing.T) {
		got := Hello("JoniBoi", "spanish")
		want := "Hola, JoniBoi"

		assertCorrectMessage(t, got, want)
	})
}
