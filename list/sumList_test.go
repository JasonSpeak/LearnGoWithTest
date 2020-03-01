package list

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Sum of 5 nums", func(t *testing.T) {
		array := [5]int{1, 2, 3, 4, 5}

		got := sum(array)
		want := 15
		if got != want {
			t.Errorf("got %d but want %d by array %v", got, want, array)
		}
	})

	t.Run("sum of any size nums:", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6}
		got := SumOfSlice(array)
		want := 21
		if got != want {
			t.Errorf("Got %d but want %d with %v", got, want, array)
		}
	})

	t.Run("Sums of arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
