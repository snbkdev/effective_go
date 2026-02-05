package _3_efficient

import (
	"fmt"
	"testing"
)

var result string

func BenchmarkToString(b *testing.B) {
	scenarios := []struct {
		desc        string
		totalPeople int
	}{
		{
			desc:        "1 person",
			totalPeople: 1,
		},
		{
			desc:        "10 people",
			totalPeople: 10,
		},
		{
			desc:        "100 people",
			totalPeople: 100,
		},
	}

	for _, s := range scenarios {
		scenario := s
		b.Run(scenario.desc, func(b *testing.B) {
			people := make([]Person, scenario.totalPeople)
			for x := 0; x < scenario.totalPeople; x++ {
				people[x] = Person{
					ID:   x,
					Name: fmt.Sprintf("test %d", x),
				}
			}

			for i := 0; i < b.N; i++ {
				result = toString(people)
			}
		})
	}
}

func toString(people []Person) string {
	var out string
	for _, person := range people {
		out += fmt.Sprintf("ID: %d\nName: %s", person.ID, person.Name)
	}
	return out
}

type Person struct {
	ID   int
	Name string
}