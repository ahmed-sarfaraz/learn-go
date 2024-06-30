package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		gotThis := Hello("Chris")
		wantThis := "Hello, Chris"
	
		if gotThis != wantThis {
			t.Errorf("got %q want %q", gotThis, wantThis)
		}
	})

	t.Run("say 'Hello world' when no parameter passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}