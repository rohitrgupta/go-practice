package integers

import "testing"

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
	t.Run("Test EvenOdd", func(t *testing.T) {
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


