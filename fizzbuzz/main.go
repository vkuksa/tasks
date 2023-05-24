package main

import (
	"errors"
	"fizzbuzz/solver"
	"fizzbuzz/solver/bruteforce"
	"fizzbuzz/solver/parametrised"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	approach = flag.String("a", "brute", "Choose approach to solve a problem. Available options: [brute, concurrent, parametrised]. Default: brute")
	begin    = flag.Int("b", 0, "Number to perform fizzbuzz from. Default: 0")
	end      = flag.Int("e", 100, "Number to perform fizzbuzz till. Default: 100")
)

func main() {
	flag.Parse()

	s, err := createSolver()
	if err != nil {
		log.Fatal(err)
	}

	var results []string
	if results, err = s.Solve(); err != nil {
		log.Fatal(err)
	}

	io.WriteString(os.Stdout, strings.Join(results, ", "))
}

func createSolver() (solver.Solver, error) {
	switch *approach {
	case "brute":
		o := bruteforce.Options{N: *end}
		return bruteforce.NewSolver(o), nil
	case "parametrised":
		t1 := parametrised.NewTerm(func(v int) bool { return v%3 == 0 }, "Fizz")
		t2 := parametrised.NewTerm(func(v int) bool { return v%5 == 0 }, "Buzz")
		t3 := parametrised.NewTerm(func(v int) bool { return v%15 == 0 }, "FizzBuzz")

		terms := []*parametrised.Term[int]{t1, t2, t3}

		o, err := parametrised.NewOptions(*begin, *end, terms)
		if err != nil {
			return nil, fmt.Errorf("createSolver: %w", err)
		}
		return parametrised.NewSolver(o), nil
	default:
		return nil, errors.New("createSolver: unknown options type provided")
	}
}
