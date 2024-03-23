package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}	


func ExampleRepeat() {
	got := Repeat("b", 6)
	fmt.Println(got)
	// Output: bbbbbb
}


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
