package main

import "testing"

func assertCorrectMessage(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
func TestHello(t *testing.T) {

	t.Run("say hello to mary", func(t *testing.T) {
		got := Hello("Mary", "")
		want := "Hello, Mary!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)

	})

	t.Run("spanish verson", func(t *testing.T) {
		got := Hello("Max", "Spanish")
		want := "Hola, Max!"

		assertCorrectMessage(t, got, want)
	})

}
