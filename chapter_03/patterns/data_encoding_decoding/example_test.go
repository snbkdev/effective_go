package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func BenchmarkExample(b *testing.B) {
	total := 10000
	users := make([]*User, total)
	for x := 0; x < total; x++ {
		users[x] = &User{
			ID:      x,
			Name:    fmt.Sprintf("Test user %d", x),
			Email:   fmt.Sprintf("%d@example.com", x),
			City:    fmt.Sprintf("Test City %d", x),
			Country: "Australia",
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, user := range users {
			encoded, _ := json.Marshal(user)

			result := &User{}
			_ = json.Unmarshal(encoded, result)

			assert.Equal(b, result.ID, user.ID)
		}
	}
}

func BenchmarkFixed(b *testing.B) {
	total := 10000
	users := make([]*UserFixed, total)
	for x := 0; x < total; x++ {
		users[x] = &UserFixed{
			ID:      int64(x),
			Name:    fmt.Sprintf("Test user %d", x),
			Email:   fmt.Sprintf("%d@example.com", x),
			City:    fmt.Sprintf("Test City %d", x),
			Country: "Australia",
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, user := range users {
			encoded, _ := proto.Marshal(user)

			result := &UserFixed{}
			_ = proto.Unmarshal(encoded, result)

			assert.Equal(b, result.ID, user.ID)
		}
	}
}