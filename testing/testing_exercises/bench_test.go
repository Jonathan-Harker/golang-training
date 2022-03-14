package testing_exercises

import (
	"testing"
)

const number = 100

func BenchmarkSprintfTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printNumberSprint(number)
	}
}

func BenchmarkFormatIntTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printNumberFormatInt(number)
	}
}

func BenchmarkItoaTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printNumberItoa(number)
	}
}
