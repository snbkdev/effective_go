package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPartitionedMap_basics(t *testing.T) {
	pMap := NewPartitionedMap(10)

	result, resultFound := pMap.Get("foo")
	assert.Nil(t, result)
	assert.False(t, resultFound)

	pMap.Set("foo", "bar")

	result, resultFound = pMap.Get("foo")
	assert.Equal(t, "bar", result)
	assert.True(t, resultFound)
}