package Repeat

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Test repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("c", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("q", 3)
	fmt.Println(result)
	//Output: qqq
}
