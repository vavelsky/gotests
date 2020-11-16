package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to other people", func(t *testing.T) {
		got := Hello("Pawel", "")
		want := "Hello Pawel"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello world! when no name provided'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean", "French")
		want := "Bonjour Jean"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Antonio", "Italian")
		want := "Ciao Antonio"

		assertCorrectMessage(t, got, want)
	})
}
