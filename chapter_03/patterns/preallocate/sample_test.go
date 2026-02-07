package main

import (
	"testing"
)

func BenchmarkExample(b *testing.B) {
	total := 1000000

	for i := 0; i < b.N; i++ {
		var data []string

		for x := 0; x < total; x++ {
			data = append(data, "x")
		}
	}
}

func BenchmarkFixed(b *testing.B) {
	total := 1000000

	for i := 0; i < b.N; i++ {
		data := make([]string, total)

		for x := 0; x < total; x++ {
			data[x] = "x"
		}
	}
}