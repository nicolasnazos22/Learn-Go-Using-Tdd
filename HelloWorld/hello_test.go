package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

func TestHelloEnglish(t *testing.T) {
	t.Run("telling hello to people", func(t *testing.T) {
		got := helloWorld("English", "Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
}

func TestHelloSpanish(t *testing.T) {
	t.Run("telling hello in Spanish", func(t *testing.T) {
		got := helloWorld("Spanish", "Juan")
		want := "Hola, Juan"
		assertCorrectMessage(t, got, want)
	})
}
func TestHelloFrench(t *testing.T) {
	t.Run("telling hello in French", func(t *testing.T) {
		got := helloWorld("French", "Agnese")
		want := "Bonjour, Agnese"
		assertCorrectMessage(t, got, want)
	})
}

func TestHello3(t *testing.T) {
	t.Run("telling hello in another language does not work", func(t *testing.T) {
		got := helloWorld("vietnamese", "John")
		want := "Hello, John"
		assertCorrectMessage(t, got, want)
	})
}
