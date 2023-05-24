// Implements a parametrised solution of FizzBuzz problem
// It currently supports only Integer subtypes, but in the future, if golang will be supporting of
// constraints that detect types that can be added or substracted, those may be added

package parametrised

import (
	"errors"
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

type Priority int

const (
	HighPriority = iota
	MediumPriority
	LowPriority
)

type TermFunc[V constraints.Integer] func(v V) bool

// Defines a term to compare agains.
type Term[V constraints.Integer] struct {
	// Predicate we use to compare against
	pred TermFunc[V]
	// String we return on successfull predicate hit
	str string
	// Priority of term
	priority Priority
	// This value signifies, whether check of terms should be proceeded if this term encountered
	interrupting bool
}

// Builder of a new term
func NewTerm[V constraints.Integer](p TermFunc[V], s string, pr Priority, i bool) *Term[V] {
	return &Term[V]{pred: p, str: s, priority: pr, interrupting: i}
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
// !!! Important it's a caller's responisibility to provide terms in asuitable way
// Returns error, If begin is larger than end
func NewOptions[V constraints.Integer](b, e V, ts []*Term[V]) (*Options[V], error) {
	if b > e {
		return nil, errors.New("parametrised newoptions: begin is set to be larger than end")
	}

	sort.Slice(ts, func(i, j int) bool {
		return ts[i].priority < ts[j].priority
	})

	return &Options[V]{b: b, e: e, ts: ts}, nil
}

type Solver[V constraints.Integer] struct {
	o *Options[V]
}

func NewSolver[V constraints.Integer](o *Options[V]) *Solver[V] {
	return &Solver[V]{o: o}
}

// Solving function
// Returns no errors
func (s *Solver[V]) Solve() ([]string, error) {
	var res = make([]string, 0)

	for i := s.o.b; i <= s.o.e; i++ {

		// We should know whether at least one of our terms is hit.
		hit := false
		for _, term := range s.o.ts {

			if term.pred(i) {
				res = append(res, term.str)
				hit = true

				// If term does not allow any other terms to execute after
				if term.interrupting {
					break
				}
			}
		}

		// If none were hit - we add just a string value
		if !hit {
			res = append(res, fmt.Sprint(i))
		}
	}

	return res, nil
}
