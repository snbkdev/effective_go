package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func CleanExample(person Person) string {
	return fmt.Sprintf("ID: %d\n Name: %s", person.ID, person.Name)
}

type Person struct {
	ID int64
	Name string
}

func FastExample(buffer *bytes.Buffer, person Person) string {
	buffer.Reset()

	_, _ = buffer.WriteString("ID: ")
	_, _ = buffer.WriteString(strconv.FormatInt(int64(person.ID), 10))
	_, _ = buffer.WriteString("\nName: ")
	_, _ = buffer.WriteString(person.Name)

	return buffer.String()
}