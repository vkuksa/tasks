package solver

// Interface representing a solver of a problem in a system
type FizzBuzzSolver interface {
	// Solve generates a solution of a problem
	// Takes n as an integer as per description
	// Solve returns an error if one occured
	Solve(n int) ([]string, error)
}
