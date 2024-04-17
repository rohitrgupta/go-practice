package vovel_count

import (
	"fmt"
	"testing"
)


func TestVovelCount(t *testing.T) {
	got := VovelCount("hello")
	expected := 2

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func ExampleVovelCount() {
	got := VovelCount("rohitrgupta")
	fmt.Println(got)
	// Output: 4
}

