package main

import (
	"testing"
)

func TestHello(t*testing.T){
	t.Run("saying hello to people", func(t*testing.T){
		input := "Alice"
		got := Hello(input)
		want := "Hello, Alice"
	
		if got != want{
			t.Errorf("Test Hello Failed got %q; want %q\n", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t*testing.T){
		got := Hello("")
		want := "Hello, World"
	
		if got != want{
			t.Errorf("Test Hello Failed got %q; want %q\n", got, want)
		}
	})
}