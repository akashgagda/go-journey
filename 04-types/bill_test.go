package main

import (
	"math"
	"testing"
)

func TestBillTotal(t *testing.T) {
	got := billTotal(12.5, 2, false)
	want := 25.0
	if math.Abs(got-want) > 1e-9 {
		t.Errorf("got %g, want %.2f", got, want)
	}
	got2 := billTotal(12.5, 2, true)
	want2 := 29.5
	if math.Abs(got2-want2) > 1e-9 {
		t.Errorf("got2 %g, want2 %.2f", got2, want2)
	}
	got3 := billTotal(10, 0, true)
	want3 := 0.0
	if math.Abs(got3-want3) > 1e-9 {
		t.Errorf("got3 %g, want3 %.2f", got3, want3)
	}

}
