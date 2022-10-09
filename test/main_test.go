package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello to mary", func(t *testing.T) {
		got := Hello("Mary")
		want := "Hello, Mary!"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("say hello world when name is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

}
