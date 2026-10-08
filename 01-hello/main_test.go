package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("akash")
	want := "hello akash"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
