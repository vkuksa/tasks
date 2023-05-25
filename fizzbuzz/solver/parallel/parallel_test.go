package parallel_test

import (
	"fmt"
	"runtime"
	"sort"
	"testing"

	p "github.com/vkuksa/tasks/fizzbuzz/solver/parallel"

	"github.com/stretchr/testify/assert"
)

func TestParallel_Solve(t *testing.T) {
	testCases := []struct {
		num      int
		expected []string
	}{
		{25, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17",
			"Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz"}},
		{10, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{1, []string{"1"}},
		{0, []string{}},
	}

	for _, nWorkers := range []int{1, 2, 4, 8, 16} {
		t.Run(fmt.Sprintf("workers_%d", nWorkers), func(t *testing.T) {
			for _, tc := range testCases {
				// Create a new Parallel solver
				o := p.NewOptions(nWorkers)
				solver := p.NewSolver(o)

				// Obrain results
				results, err := solver.Solve(tc.num)
				assert.NoError(t, err, "Error solving FizzBuzz for num=%d: %v", tc.num, err)

				// The order of content of parallel solver would be unspecified
				// So we assert if they have same content
				assert.True(t, haveSameElements(results, tc.expected), "For num=%d, expected output:\n%s\n\nActual output:\n%s\n", tc.num, tc.expected, results)
			}
		})
	}
}

func haveSameElements(first, second []string) bool {
	// Sort the slices
	sort.Strings(first)
	sort.Strings(second)

	// Check if the sorted slices are equal
	if len(second) != len(first) {
		return false
	}

	// Compare sorted content
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func BenchmarkParallel(b *testing.B) {
	// Create a new BruteForce solver with the buffer as the output writer
	o := p.NewOptions(runtime.NumCPU())
	solver := p.NewSolver(o)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = solver.Solve(100)
	}
}
