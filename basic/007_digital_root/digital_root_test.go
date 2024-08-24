package digital_root

import (
	"fmt"
	"testing"
)


func TestDigitalRoot(t *testing.T) {
	got := DigitalRoot(16)
	expected := 7

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func ExampleDigitalRoot() {
	got := DigitalRoot(195)
	fmt.Println(got)
	// Output: 6
}
