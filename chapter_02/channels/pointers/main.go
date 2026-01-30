package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	john := &Person{Name:"Bob"}

	peopleCh := make(chan *Person, 1)
	peopleCh <- john

	john.Name = "Paul"

	result := <- peopleCh
	assert.Equal(t, "Bob", result.Name)
}

type Person struct {
	Name string
}