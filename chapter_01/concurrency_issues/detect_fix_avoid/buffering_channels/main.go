package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelStats(t *testing.T) {
	bufferedCh := make(chan struct{}, 10)

	bufferedCh <- struct{}{}
	bufferedCh <- struct{}{}
	bufferedCh <- struct{}{}

	assert.Equal(t, 3, len(bufferedCh))
}