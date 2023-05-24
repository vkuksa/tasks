package parametrised

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type TermFunc[V constraints.Integer] func(v V) bool

// Defines a term to compare agains.
type Term[V constraints.Integer] struct {
	// Predicate we use to compare against
	pred TermFunc[V]
	// String we return on successfull predicate hit
	s string
}

// Builder of a new term
func NewTerm[V constraints.Integer](p TermFunc[V], s string) *Term[V] {
	return &Term[V]{pred: p, s: s}
}

// Defines options of execution
type Options[V constraints.Integer] struct {
	// Beginning value that will be assigned as the start of a loop
	b V
	// Ending value that will be assigned as the end of a loop
	e V
	// Terms to compare values against
	ts []*Term[V]
}

// Builder of options
func NewOptions[V constraints.Integer](b, e V, ts []*Term[V]) (*Options[V], error) {
	// If begin is larger than end - return an error
	if b > e {
		return nil, errors.New("parametrised newoptions: begin is set to be larger than end")
	}

	return &Options[V]{b: b, e: e, ts: ts}, nil
}

type Solver[V constraints.Integer] struct {
	o *Options[V]
}

func NewSolver[V constraints.Integer](o *Options[V]) *Solver[V] {
	return &Solver[V]{o: o}
}

// Solving function
func (s *Solver[V]) Solve() ([]string, error) {
	var res = make([]string, 0)

	for i := s.o.b; i <= s.o.e; i++ {

		// We should know whether at least one of our terms is hit.
		hit := false
		for _, t := range s.o.ts {
			if t.pred(i) {
				res = append(res, t.s)
				hit = true
			}
		}

		// If none were hit - we add just a string value
		if !hit {
			res = append(res, fmt.Sprint(i))

		}
	}

	return res, nil
}
