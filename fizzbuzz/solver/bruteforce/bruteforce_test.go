package bruteforce_test

import (
	"strings"
	"testing"

	bf "github.com/vkuksa/tasks/fizzbuzz/solver/bruteforce"

	"github.com/stretchr/testify/assert"
)

func TestBrutForce_Solve(t *testing.T) {
	testCases := []struct {
		num      int
		expected string
	}{
		{25, "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz1617Fizz19BuzzFizz2223FizzBuzz"},
		{10, "12Fizz4BuzzFizz78FizzBuzz"},
		{5, "12Fizz4Buzz"},
		{1, "1"},
		{0, ""},
	}

	for _, tc := range testCases {
		// Create a new BruteForce solver
		solver := bf.NewSolver()

		// Call the Solve method to obrain results
		results, err := solver.Solve(tc.num)
		assert.NoError(t, err, "Error solving FizzBuzz for num=%d: %v", tc.num, err)

		// Compare the actual and expected output
		output := strings.Join(results, "")
		assert.Equal(t, output, tc.expected, "For num=%d, expected output:\n%s\n\nActual output:\n%s\n", tc.num, tc.expected, output)
	}
}

func BenchmarkBruteForce_Solve(b *testing.B) {
	// Create a new BruteForce solver with the buffer as the output writer
	solver := bf.NewSolver()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = solver.Solve(100)
	}
}
