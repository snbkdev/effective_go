package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsePooledAccount(t *testing.T) {
	assert.Equal(t, 5, usePooledAccount())
	assert.Equal(t, 5, usePooledAccount())
}