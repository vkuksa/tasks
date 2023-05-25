package parametrised_test

import (
	"strings"
	"testing"

	p "github.com/vkuksa/tasks/fizzbuzz/solver/parametrised"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

type testCase[V constraints.Integer] struct {
	num      V
	expected string
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
	terms := p.NewTermCollection(
		p.NewTerm(func(v V) bool { return v%3 == 0 }, "Fizz", p.LowPriority, false),
		p.NewTerm(func(v V) bool { return v%5 == 0 }, "Buzz", p.LowPriority, false),
		p.NewTerm(func(v V) bool { return v%15 == 0 }, "FizzBuzz", p.HighPriority, true),
	)

	// Test cases
	for _, tc := range tc {
		opt := p.NewOptions(1, terms)
		solver := p.NewSolver(opt)

		result, err := solver.Solve(0)
		assert.Error(t, err)

		result, err = solver.Solve(tc.num)
		assert.NoError(t, err)

		assert.Equal(t, tc.expected, strings.Join(result, ""))
	}
}

func BenchmarkParametrised_Solve(b *testing.B) {
	terms := p.NewTermCollection(
		p.NewTerm(func(v int) bool { return v%3 == 0 }, "Fizz", p.LowPriority, false),
		p.NewTerm(func(v int) bool { return v%5 == 0 }, "Buzz", p.LowPriority, false),
		p.NewTerm(func(v int) bool { return v%15 == 0 }, "FizzBuzz", p.HighPriority, true),
	)

	opt := p.NewOptions(1, terms)
	solver := p.NewSolver(opt)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = solver.Solve(100)
	}
}
