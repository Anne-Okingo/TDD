package main

import (
	"testing"
)

func TestHello(t*testing.T){
	t.Run("saying hello to people", func(t*testing.T){
		input := "Alice"
		got := Hello(input, "")
		want := "Hello, Alice"
	

		assertMessage(t,got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t*testing.T){
		got := Hello("","")
		want := "Hello, World"
	
		assertMessage(t,got,want)
	})
	t.Run("in German", func(t*testing.T){
		got := Hello("Temo", "Deutsch")
		want := "Hallo, Temo"
		assertMessage(t,got,want)
	})
}

func assertMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want{
		t.Errorf("Testing Hello failed got %q; want %q\n",got, want)
	}
}