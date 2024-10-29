package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("say for this example", func(t *testing.T) {
		input := "a"
		num := 7
		got := Repeat(input, num)
		want := "aaaaaaa"

		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Test Repeat FAilled : got: %s, want:  %s\n", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}

// If ever your code changes so that the example is no longer valid, your build will fail.
func ExampleRepeat() {
	result := Repeat("woo", 3)

	fmt.Println(result)
	//Output: woowoowoo
}
