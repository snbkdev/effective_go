package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheWithMutex(t *testing.T) {
	cache := newCacheUsingMutex()

	cache.Set("my-key", Person{Name: "King"})

	result := cache.Get("my-key")

	expected := Person{Name: "King"}
	assert.Equal(t, expected, result)
}