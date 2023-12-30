package main

import (
	"bytes"
	"testing"
)

func TestWordCount(t *testing.T) {
	b := bytes.NewBufferString("An old silent pond A frog jumps into the pond— Splash! Silence again.\n")

	exp := 13
	res := count(b, false, false)
	if res != exp {
		t.Errorf("Expected %d words, got %d", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("An old silent pond\n A frog jumps into the pond\n— Splash! Silence again.\n")

	exp := 3
	res := count(b, true, false)
	if res != exp {
		t.Errorf("Expected %d lines, got %d", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("An old silent pond A frog jumps into the pond— Splash! Silence again.\n")

	exp := 72
	res := count(b, false, true)
	if res != exp {
		t.Errorf("Expected %d bytes, got %d", exp, res)
	}
}
