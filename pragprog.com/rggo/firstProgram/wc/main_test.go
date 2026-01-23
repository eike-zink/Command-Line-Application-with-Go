package main

import (
	"bytes"
	"testing"
)

// TestCountWords test the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n")

	exp := 3

	res := count(b)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
