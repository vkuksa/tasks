package bruteforce_test

import (
	"bytes"
	"testing"

	bf "fizzbuzz/solver/bruteforce"
)

func TestBrutForce_WriteToNilWriter(t *testing.T) {
	// Create New BruteForce solver
	solver := bf.NewSolver()

	// Call the Solve method with the specific number
	err := solver.Solve(nil, 1)
	if err == nil {
		t.Fatalf("Expected error for writer equals nil")
	}

}

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

	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Create a new BruteForce solver with the buffer as the output writer
	solver := bf.NewSolver()

	for _, tc := range testCases {
		// Call the Solve method with the specific number
		err := solver.Solve(&buf, tc.num)
		if err != nil {
			t.Fatalf("Error solving FizzBuzz for num=%d: %v", tc.num, err)
		}

		// Get the captured output
		output := buf.String()

		// Compare the actual and expected output
		if output != tc.expected {
			t.Fatalf("For num=%d, expected output:\n%s\n\nActual output:\n%s\n", tc.num, tc.expected, output)
		}

		// Reset the buffer for a next test execution
		buf.Reset()
	}
}

func BenchmarkBruteForce_Solve(b *testing.B) {
	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Create a new BruteForce solver with the buffer as the output writer
	solver := bf.NewSolver()

	// Reset the buffer and benchmark the Solve method
	buf.Reset()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = solver.Solve(&buf, 100)
	}
}
