package main

import (
	"fmt"
	"testing"
)

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Finish the solution so that it returns the sum of all the multiples of 3 or 5 below the number passed in.

// Note: If the number is a multiple of both 3 and 5, only count it once.const

 

func TestMultiple3And5(t *testing.T) {
	got := Multiple3And5(10)
	expected := 23
	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}


func ExampleMultiple3And5() {
	got := Multiple3And5(11)
	fmt.Println(got)
	// Output: 33
}
