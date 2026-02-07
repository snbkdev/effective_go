package main

import (
	"crypto/md5"
	"testing"
)

func sign(msg string) string {
	hash := md5.New()

	result := hash.Sum([]byte(msg))
	return string(result)
}

func BenchmarkExample(b *testing.B) {
	input := "hello universe!!!!"

	for i := 0; i < b.N; i++ {
		sign(input)
	}
}