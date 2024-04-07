package main

import "testing"

func TestUnpacker(t *testing.T) {
	got, _ := Unpack("a4bc2d5e3")
	want := "aaaabccdddddeee"
	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}

func TestError(t *testing.T) {
	_, err := Unpack("a4bc2d5e3")
	if err == nil {
		t.Errorf("Expect error but it isn't")
	}
}