package bruteforce

import (
	"strconv"
)

type Options struct {
	// Specifies the end of execution, when incrementing loop index reaches this value
	N int
}

type Solver struct {
	o *Options
}

func NewSolver(o *Options) *Solver {
	return &Solver{o: o}
}

// Classic solution of FizzBuzz
func (s *Solver) Solve() ([]string, error) {
	var res = make([]string, 0, s.o.N)

	for i := 1; i <= s.o.N; i++ {
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
