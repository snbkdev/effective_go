package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var simpleTotal int64

func BenchmarkSimpleAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < 1000000; x++ {
			simpleTotal += int64(x)
		}

		fmt.Printf("total: %d\n", simpleTotal)
	}
}

var atomicTotal int64

func BenchmarkAtomicAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < 1000000; x++ {
			atomic.AddInt64(&atomicTotal, int64(x))
		}

		fmt.Printf("total: %d\n", atomicTotal)
	}
}

var mutexTotal int64
var mutex sync.Mutex

func BenchmarkMutexAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < 1000000; x++ {
			mutex.Lock()
			mutexTotal += int64(x)
			mutex.Unlock()
		}

		fmt.Printf("total: %d\n", mutexTotal)
	}
}