package main

import "testing"

func TestConcurrency(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping concurrent test because of short mode")
	}
}