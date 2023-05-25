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

// Defines a set of priorities for terms
// 0 corresponds with highest priority
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

// Defines a storage of Terms
type TermCollection[V constraints.Integer] []*Term[V]

func NewTermCollection[V constraints.Integer](terms ...*Term[V]) TermCollection[V] {
	t := make(TermCollection[V], 0, len(terms))
	for _, term := range terms {
		t = append(t, term)
	}
	return t
}

func (t *TermCollection[V]) SortByPriority() {
	sort.Slice(*t, func(i, j int) bool {
		return (*t)[i].priority < (*t)[j].priority
	})
}

// Defines options of execution
type Options[V constraints.Integer] struct {
	// Beginning value that will be assigned as the start of a loop
	b V
	// Terms to compare values against
	ts TermCollection[V]
}

// Builder of options
// Performs sorting of terms by specified
// Returns error, If begin is larger than end
func NewOptions[V constraints.Integer](b V, ts TermCollection[V]) *Options[V] {
	ts.SortByPriority()
	return &Options[V]{b: b, ts: ts}
}

type Solver[V constraints.Integer] struct {
	o *Options[V]
}

// Creates new Solver instance
func NewSolver[V constraints.Integer](o *Options[V]) *Solver[V] {
	return &Solver[V]{o: o}
}

// Solve takes n as input, to generate values till it reached starting from 1
// Refer to https://leetcode.com/problems/fizz-buzz/description/
// Solve returns an error, if value for begin, that was in options, is larger than n
func (s *Solver[V]) Solve(n V) ([]string, error) {
	if n < s.o.b {
		return []string{}, errors.New("Provided begin value is larger than n")
	}
	var res = make([]string, 0)

	for i := s.o.b; i <= n; i++ {

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
