package main

import "testing"

func BenchmarkBasics(b *testing.B) {
	// build input definition and perform initialization

	// remove init time from measurement

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		//
	}
}