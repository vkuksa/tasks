package parallel

import (
	"strconv"
)

type BruteForce struct{}

func NewSolver() *BruteForce {
	return &BruteForce{}
}

func (bf *BruteForce) Solve(n int) ([]string, error) {
	var res = make([]string, 0, n)

	for i := 1; i <= n; i++ {
		switch {
		// Deliberately omitting fallthrough here to increase readability
		case i%15 == 0:
			res = append(res, "FizzBuzz")
		case i%3 == 0:
			res = append(res, "Fizz")
		case i%5 == 0:
			res = append(res, "Buzz")
		default:
			res = append(res, strconv.Itoa(i))
		}
	}

	return res, nil
}
