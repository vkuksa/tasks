package solver

// Interface representing a solver of a problem in a system
type Solver interface {
	// Solve generates a solution of a problem
	// Solve returns an error if one occured
	Solve() ([]string, error)
}
