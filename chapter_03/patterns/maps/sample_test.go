package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkMapString(b *testing.B) {
	data := map[string]string{}
	for x := 100000; x < 1000000; x++ {
		data[strconv.Itoa(x)] = "X"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := data[strconv.Itoa(500000)]
		require.Equal(b, "X", result)
	}
}

func BenchmarkMapInt(b *testing.B) {
	data := map[int]string{}
	for x := 100000; x < 1000000; x++ {
		data[x] = "X"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := data[500000]
		require.Equal(b, "X", result)
	}
}