package main

import "testing"


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
