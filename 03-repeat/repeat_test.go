package repeat

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("ab", 3)
	want := "ababab"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	got1 := Repeat("x", 0)
	want1 := ""
	if got1 != want1 {
		t.Errorf("got1 %s, want1 %s", got1, want1)

	}

	got2 := Repeat("go", 2)
	want2 := "gogo"
	if got2 != want2 {
		t.Errorf("got2 %s, want2 %s", got2, want2)
	}
}
