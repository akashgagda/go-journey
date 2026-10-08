package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(2, 2)

	got1 := Add(99, 1)

	want := 4

	want1 := 100

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	if got1 != want1 {
		t.Errorf("got1 %d, want1 %d", got1, want1)
	}
}
