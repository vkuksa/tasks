package parametrised_test

import (
	"fizzbuzz/solver/parametrised"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

type testCase[V constraints.Integer] struct {
	num      V
	expected string
}

func TestParametrisedSolver_CheckOptionsError(t *testing.T) {
	_, err := parametrised.NewOptions(1, 0, []*parametrised.Term[int]{})
	assert.Error(t, err)
}

func TestParametrisedSolver_ints(t *testing.T) {
	testCasesInt64 := []testCase[int64]{
		{25, "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz1617Fizz19BuzzFizz2223FizzBuzz"},
		{10, "12Fizz4BuzzFizz78FizzBuzz"},
		{5, "12Fizz4Buzz"},
		{1, "1"},
	}
	checkTestCases(t, testCasesInt64)

	testCasesUint32 := []testCase[uint32]{
		{25, "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz1617Fizz19BuzzFizz2223FizzBuzz"},
		{10, "12Fizz4BuzzFizz78FizzBuzz"},
		{5, "12Fizz4Buzz"},
		{1, "1"},
	}
	checkTestCases(t, testCasesUint32)
}

func checkTestCases[V constraints.Integer](t testing.TB, tc []testCase[V]) {
	// Test cases
	for _, tc := range tc {
		terms := []*parametrised.Term[V]{
			parametrised.NewTerm(func(v V) bool {
				return v%3 == 0
			}, "Fizz"),
			parametrised.NewTerm(func(v V) bool {
				return v%5 == 0
			}, "Buzz"),
		}

		opt, err := parametrised.NewOptions(1, tc.num, terms)
		assert.NoError(t, err)

		solver := parametrised.NewSolver(opt)
		result, err := solver.Solve()
		assert.NoError(t, err)

		assert.Equal(t, tc.expected, strings.Join(result, ""))
	}
}

func BenchmarkBruteForce_Solve(b *testing.B) {
	terms := []*parametrised.Term[int]{
		parametrised.NewTerm(func(v int) bool {
			return v%3 == 0
		}, "Fizz"),
		parametrised.NewTerm(func(v int) bool {
			return v%5 == 0
		}, "Buzz"),
	}

	opt, _ := parametrised.NewOptions(1, 100, terms)
	solver := parametrised.NewSolver(opt)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = solver.Solve()
	}
}
