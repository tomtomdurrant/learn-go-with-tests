package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("Got %d, want %d, values %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{3, 4}

	got := SumAll(slice1, slice2)
	want := []int{3, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	CheckSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %d, want %d", got, want)
		}
	}

	t.Run("sum of some slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{3, 4}

		got := SumAllTails(slice1, slice2)
		want := []int{2, 4}

		CheckSums(t, got, want)
	})

	t.Run("safely handles a 0 length slice", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{1, 2, 3}

		got := SumAllTails(slice1, slice2)
		want := []int{0, 5}

		CheckSums(t, got, want)
	})
}
