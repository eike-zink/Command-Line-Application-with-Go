package main

import (
	"bytes"
	"testing"
)

// TestCountWords test the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n")

	exp := 3

	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("line1\n line2\n line3\n")

	exp := 3

	res := count(b, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
