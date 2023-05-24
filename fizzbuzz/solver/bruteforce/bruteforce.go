package bruteforce

import (
	"strconv"
)

// In this variation only calibrating from 1 to n is available
type Options struct {
	N int
}

type BruteForce struct {
	o Options
}

func NewSolver(opt Options) *BruteForce {
	return &BruteForce{o: opt}
}

func (bf *BruteForce) Solve() ([]string, error) {
	var res = make([]string, 0, bf.o.N)

	for i := 1; i <= bf.o.N; i++ {
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
