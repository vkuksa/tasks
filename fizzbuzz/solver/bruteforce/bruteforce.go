// Implements classic solution of FizzBuzz
package bruteforce

import (
	"strconv"
)

type Solver struct{}

func NewSolver() *Solver {
	return &Solver{}
}

// Solve takes n as input, to generate values till it reached starting from 1
// Refer to https://leetcode.com/problems/fizz-buzz/description/
// Solve does not return any errors
func (s *Solver) Solve(n int) ([]string, error) {
	var res = make([]string, 0, n)

	for i := 1; i <= n; i++ {
		switch {
		// Deliberately omitting fallthrough approach here to increase readability
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
