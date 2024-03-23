package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func areEq(a, b []int) bool {
	if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func TestArrayDiff(t *testing.T) {
	t.Run("difference of arrays", func(t *testing.T) {
		input := []int{1, 2, 2, 2, 3}
		other := []int{2}
		want := []int{1, 3}
		got := ArrayDiff(input, other)
		if !areEq(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	
	})
}