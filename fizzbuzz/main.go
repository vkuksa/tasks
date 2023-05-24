package main

import (
	"errors"
	"fizzbuzz/solver"
	"fizzbuzz/solver/bruteforce"
	"flag"
	"log"
	"os"
)

var (
	approach = flag.String("a", "brute", "Choose approach to solve a problem. Available options: [brute, concurrent]. Default: brute")
	number   = flag.Int("n", 100, "Number to perform fizzbuzz against. Default: 100")
)

func main() {
	flag.Parse()

	s, err := createSolver()
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Solve(os.Stdout, *number); err != nil {
		log.Fatal(err)
	}
}

func createSolver() (solver.Solver, error) {
	switch *approach {
	case "brute":
		return bruteforce.NewSolver(), nil
	default:
		return nil, errors.New("createSolver: unknown options type provided")
	}
}
