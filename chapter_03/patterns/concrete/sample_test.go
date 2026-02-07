package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func Sign(writer io.Writer) error {
	_, err := writer.Write([]byte("Signed by Me!"))
	return err
}

func BenchmarkSign(b *testing.B) {
	in := bytes.NewBufferString("test")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := Sign(in)
		require.NoError(b, err)
	}
}

func SignFixed(writer *bytes.Buffer) error {
	_, err := writer.Write([]byte("Signed by Me!"))
	return err
}

func BenchmarkSignFixed(b *testing.B) {
	in := bytes.NewBufferString("test")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := SignFixed(in)
		require.NoError(b, err)
	}
}