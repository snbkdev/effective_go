package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

var result string

func BenchmarkExample(b *testing.B) {
	start := time.Now()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = fmt.Sprintf("%d ns", time.Since(start).Nanoseconds())
	}
}

func BenchmarkStrconv(b *testing.B) {
	start := time.Now()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = strconv.FormatInt(time.Since(start).Nanoseconds(), 10) + " ns"
	}
}

func BenchmarkAppend(b *testing.B) {
	start := time.Now()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var dest []byte
		dest = strconv.AppendInt(dest, time.Since(start).Nanoseconds(), 10)
		dest = append([]byte(" ns"))
		result = string(dest)
	}
}

func BenchmarkBuilder(b *testing.B) {
	start := time.Now()

	b.ResetTimer()

	builder := &strings.Builder{}

	for i := 0; i < b.N; i++ {
		builder.WriteString(strconv.FormatInt(time.Since(start).Nanoseconds(), 10))
		builder.WriteString(" ns")
		result = string(builder.String())

		builder.Reset()
	}
}