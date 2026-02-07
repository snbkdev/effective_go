package main

import (
	"regexp"
	"testing"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var result bool

func BenchmarkExampleRegexOnlyBadInput(b *testing.B) {
	input := ""

	for i := 0; i < b.N; i++ {
		result = isEmail(input)
	}
}

func BenchmarkExampleRegexOnlyGoodInput(b *testing.B) {
	input := "me@home.com"

	for i := 0; i < b.N; i++ {
		result = isEmail(input)
	}
}

func isEmail(in string) bool {
	return emailRegex.MatchString(in)
}

func BenchmarkExampleFastFailBadInput(b *testing.B) {
	input := ""

	for i := 0; i < b.N; i++ {
		result = isEmailFastFail(input)
	}
}

func BenchmarkExampleFastFailGoodInput(b *testing.B) {
	input := "me@home.com"

	for i := 0; i < b.N; i++ {
		result = isEmailFastFail(input)
	}
}

func isEmailFastFail(in string) bool {
	if in == "" {
		return false
	}

	return emailRegex.MatchString(in)
}