package main

import (
	"sync"
	"testing"
)

func BenchmarkExample(b *testing.B) {
	mutex := &sync.Mutex{}

	for i := 0; i < b.N; i++ {
		doWorkWithDefer(mutex)
	}
}

func doWorkWithDefer(mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()

	// do other things
}

func BenchmarkFixed(b *testing.B) {
	mutex := &sync.Mutex{}

	for i := 0; i < b.N; i++ {
		doWorkWithoutDefer(mutex)
	}
}

func doWorkWithoutDefer(mutex *sync.Mutex) {
	mutex.Lock()

	// do other things

	mutex.Unlock()
}