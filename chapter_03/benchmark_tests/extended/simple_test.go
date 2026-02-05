package main

import (
	"fmt"
	"testing"
)

var result string

func BenchmarkToString(b *testing.B) {
	scenarios := []struct {
		desc         string
		toStringFunc func([]Person) string
	}{
		{
			desc:         "v1",
			toStringFunc: toStringV1,
		},
		{
			desc:         "v2",
			toStringFunc: toString,
		},
	}

	totalPeople := 10000
	people := make([]Person, totalPeople)
	for x := 0; x < totalPeople; x++ {
		people[x] = Person{ID: x, Name: fmt.Sprintf("test %d", x)}
	}

	for _, s := range scenarios {
		scenario := s
		b.Run(scenario.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = scenario.toStringFunc(people)
			}
		})
	}
}

func toStringV1(people []Person) string {
	var out string
	for _, person := range people {
		out += fmt.Sprintf("ID: %d\nName: %s", person.ID, person.Name)
	}
	return out
}