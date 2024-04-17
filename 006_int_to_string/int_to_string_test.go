package int_to_string

import (
	"fmt"
	"testing"
)


func TestIntToString(t *testing.T) {
	got := IntToString(12345)
	expected := "12345"

	if got != expected {
		t.Errorf("expected %s but got %s", expected, got)
	}
}

func ExampleIntToString() {
	got := IntToString(-5678)
	fmt.Println(got)
	// Output: -5678
}

