package main

import "testing"

func TestSumAll(t *testing.T) {
	//case 1
	got := sumAll([]int{1, 2, 3})
	want := 6
	//case 2
	got2 := sumAll([]int{})
	want2 := 0
	//case 3
	var s []int
	s = append(s, 4, 5)
	got3 := sumAll(s)
	want3 := 9

	if got3 != want3 {
		t.Errorf("got3 %d, want %d", got3, want3)
	}
	if got2 != want2 {
		t.Errorf("got2 %d, want2 %d", got2, want2)
	}
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
