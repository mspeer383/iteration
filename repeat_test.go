package iteration

import "testing"

var cr = 9

func TestRepeat(t *testing.T) {

	repeated := Repeat("a", cr)
	expected := "aaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %s but got %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", cr+cr)
	}
}
