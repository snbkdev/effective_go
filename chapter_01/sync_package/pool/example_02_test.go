package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsePooledAccountFixed(t *testing.T) {
	assert.Equal(t, 5, usePooledAccountFixed())
	assert.Equal(t, 5, usePooledAccountFixed())
}