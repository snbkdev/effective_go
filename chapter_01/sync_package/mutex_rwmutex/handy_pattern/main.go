package main

import (
	"errors"
	"sync"
)

type Dictionary struct {
	data []string
	mutex sync.Mutex
}

func (d *Dictionary) Get(index int) (string, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.count() >= index {
		return "", errors.New("out of bounds")
	}

	return d.data[index], nil
}

func (d *Dictionary) Count() int {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.count()
}

func (d *Dictionary) count() int {
	return len(d.data)
}