package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("B", 5)
	expected := "BBBBB"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("B", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("B", 5)
	fmt.Println(repeated)
	// Output: BBBBB
}