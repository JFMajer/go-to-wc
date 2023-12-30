package main

import (
	"bytes"
	"testing"
)

func TestWordCount(t *testing.T) {
	b := bytes.NewBufferString("An old silent pond A frog jumps into the pond— Splash! Silence again.\n")

	exp := 13
	res := count(b)
	if res != exp {
		t.Errorf("Expected %d words, got %d", exp, res)
	}
}
