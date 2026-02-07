package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		collection := &collection{}

		for x := 0; x < 10000; x++ {
			collection.add(x)
		}
		require.Equal(b, 10000, len(collection.data))
	}
}

type collection struct {
	data []int
}

func (c *collection) add(newItem int) {
	for _, thisItem := range c.data {
		if thisItem == newItem {
			return
		}
	}

	c.data = append(c.data, newItem)
}

func BenchmarkFixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withMap := &collectionWithMap{
			lookup: map[int]struct{}{},
		}

		for x := 0; x < 10000; x++ {
			withMap.add(x)
		}
		require.Equal(b, 10000, len(withMap.data))
		require.Equal(b, 10000, len(withMap.lookup))
	}
}

type collectionWithMap struct {
	data   []int
	lookup map[int]struct{}
}

func (m *collectionWithMap) add(newItem int) {
	_, found := m.lookup[newItem]
	if found {
		return
	}
	m.data = append(m.data, newItem)
	m.lookup[newItem] = struct{}{}
}