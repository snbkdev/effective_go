package main

import (
	"runtime"
	"sync/atomic"
)

type ThreadSafeSlice struct {
	data []interface{}
	reserved int64
	written int64
}

func (t *ThreadSafeSlice) Put(value interface{}) {
	index := atomic.AddInt64(&t.reserved, 1) - 1
	t.data[index] = value

	for !atomic.CompareAndSwapInt64(&t.written, index, index+1) {
		runtime.Gosched()
	}
}

func (t *ThreadSafeSlice) GetAll() []interface{} {
	currentWritten := atomic.LoadInt64(&t.written)
	return t.data[:currentWritten]
}