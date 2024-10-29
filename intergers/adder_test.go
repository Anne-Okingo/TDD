package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("say add uint intergers only", func(t *testing.T) {
		input1, input2 := 5, 4

		got := Sum(input1, input2)
		want := 9

		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Testing Sum Failed : got %d := want %d\n", got, want)
	}
}

func ExampleSum() {
	sum := Sum(4, 5)
	fmt.Println(sum)
	//Output: 9
}
