package main

import "testing"

func TestMain_AddNumbers(t *testing.T) {
	want := 3
	got := addNumbers(1, 2)

	if got != want {
		t.Errorf("addNumbers(%d, %d) => Got %d but want %d", 1, 2, got, want)
	}
}

func TestMain_SubNumbers(t *testing.T) {
	want := 3
	got := subNumbers(13, 3)

	if got != want {
		t.Errorf("subNumbers(%d, %d) => Got %d but want %d", 1, 2, got, want)
	}
}
