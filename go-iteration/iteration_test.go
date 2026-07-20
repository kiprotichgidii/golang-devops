package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("B")
	expected := "BBBBB"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("B")
	}
}