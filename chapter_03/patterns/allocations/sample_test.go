
package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Marshal(in []byte) (User, error) {
	out := User{}

	err := json.Unmarshal(in, &out)

	return out, err
}

func BenchmarkExample(b *testing.B) {
	testUser := User{ID: 666, Name: "Bob"}
	data, err := json.Marshal(testUser)
	require.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result, err := Marshal(data)

		require.NoError(b, err)
		require.Equal(b, testUser, result)
	}
}

func MarshalFixed(in []byte, user *User) error {
	return json.Unmarshal(in, user)
}

func BenchmarkFixed(b *testing.B) {
	testUser := User{ID: 666, Name: "Bob"}
	data, err := json.Marshal(testUser)
	require.NoError(b, err)

	result := &User{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := MarshalFixed(data, result)

		require.NoError(b, err)
		require.EqualValues(b, testUser, *result)
	}
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
