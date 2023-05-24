package solver

import "io"

const (
	fb = "FizzBuzz"
	f  = "Fizz"
	b  = "Buzz"
)

// Interface representing a solver of a problem in a system
type Solver interface {
	// Solve generates a solution of a problem
	// Solve returns an error if one occured
	Solve(w io.Writer, n int) error
}
