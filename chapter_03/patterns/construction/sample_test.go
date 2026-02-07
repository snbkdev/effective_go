package main

import (
	"errors"
	"fmt"
	"testing"
)

func log(debug bool, msg string) {
	if !debug {
		return
	}

	println(msg)
}

func BenchmarkExample(b *testing.B) {
	debug := false
	err := errors.New("something failed")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		log(debug, fmt.Sprintf("Error was: %s", err))
	}
}