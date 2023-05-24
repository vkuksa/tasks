package parametrised

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type TermFunc[V constraints.Integer] func(v V) bool

// Defines a term to compare agains.
type Term[V constraints.Integer] struct {
	pred TermFunc[V]
	str  string
}

func NewTerm[V constraints.Integer](pred TermFunc[V], s string) *Term[V] {
	return &Term[V]{pred: pred, str: s}
}

// Calibration from B till E
type Options[V constraints.Integer] struct {
	b V
	e V

	terms []*Term[V]
}

func NewOptions[V constraints.Integer](b, e V, t []*Term[V]) (*Options[V], error) {
	if b > e {
		return nil, errors.New("parametrised newoptions: begin is set to be larger than end")
	}

	return &Options[V]{b: b, e: e, terms: t}, nil
}

type Parametrised[V constraints.Integer] struct {
	o *Options[V]
}

func NewSolver[V constraints.Integer](opt *Options[V]) *Parametrised[V] {
	return &Parametrised[V]{o: opt}
}

// Performance of such approach is questionable
func (p *Parametrised[V]) Solve() ([]string, error) {
	var res = make([]string, 0)

	for i := p.o.b; i <= p.o.e; i++ {

		termHit := false
		for _, term := range p.o.terms {
			if term.pred(i) {
				res = append(res, term.str)
				termHit = true
			}
		}

		if !termHit {
			res = append(res, fmt.Sprint(i))

		}
	}

	return res, nil
}
