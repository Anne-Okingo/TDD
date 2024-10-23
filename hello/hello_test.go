package main

import (
	"testing"
)

func TestHello(t*testing.T){
	got := Hello("Alice")
	want := "Hello,Alice"

	if got != want{
		t.Errorf("Test Alice Failed got %q; want %q\n", got, want)
	}
}