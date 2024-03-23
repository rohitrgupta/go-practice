package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Add numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

func TestEvenOdd(t *testing.T) {
	t.Run("Test Even", func(t *testing.T) {
		got := EvenOdd(2)
		expected := "Even"
		if got != expected {
			t.Errorf("expected '%s' but got '%s'", expected, got)
		}
	})
	t.Run("Test Odd", func(t *testing.T) {
		got := EvenOdd(1)
		expected := "Odd"
		if got != expected {
			t.Errorf("expected '%s' but got '%s'", expected, got)
		}
	})
}

func TestMultiple3And5(t *testing.T) {
	got := Multiple3And5(10)
	expected := 23
	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func ExampleEvenOdd_1() {
	got := EvenOdd(5)
	fmt.Println(got)
	// Output: Odd
}

func ExampleEvenOdd_2() {
	got := EvenOdd(6)
	fmt.Println(got)
	// Output: Even
}

func ExampleMultiple3And5() {
	got := Multiple3And5(11)
	fmt.Println(got)
	// Output: 33
}
