package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheWithChannelsExample(t *testing.T) {
	cache := NewCacheUsingChannels()
	defer cache.shutdown()

	cache.Set("my-key", Person{Name:"King"})

	resultCh := cache.Get("my-key")

	select{
	case result := <- resultCh:
		expected := Person{Name:"Bob"}
		assert.Equal(t, expected, result)

	case <- time.After(1 * time.Second):
		assert.FailNow(t, "cache read time out")
	}
}