package parallel

type Options struct {
	N int
}

type Solver struct {
	o Options
}

func NewSolver(opt Options) *Solver {
	return &Solver{o: opt}
}

func (s *Solver) Solve() ([]string, error) {
	return nil, nil
}
