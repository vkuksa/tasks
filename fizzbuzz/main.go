package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/vkuksa/tasks/fizzbuzz/solver"
	bf "github.com/vkuksa/tasks/fizzbuzz/solver/bruteforce"
	pl "github.com/vkuksa/tasks/fizzbuzz/solver/parallel"
	pm "github.com/vkuksa/tasks/fizzbuzz/solver/parametrised"
)

var (
	approach = flag.String("a", "brute", "Choose approach to solve a problem. Available options: [brute, concurrent, parametrised]. Default: brute")
	begin    = flag.Int("b", 1, "Number to perform fizzbuzz from. Default: 0")
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
		o := &bf.Options{N: *end}
		return bf.NewSolver(o), nil
	case "parametrised":
		terms := make([]*pm.Term[int], 0, 3)
		terms = append(terms, pm.NewTerm(func(v int) bool { return v%3 == 0 }, "Fizz", pm.LowPriority, false))
		terms = append(terms, pm.NewTerm(func(v int) bool { return v%5 == 0 }, "Buzz", pm.LowPriority, false))
		terms = append(terms, pm.NewTerm(func(v int) bool { return v%15 == 0 }, "FizzBuzz", pm.HighPriority, true))

		o, err := pm.NewOptions(*begin, *end, terms)
		if err != nil {
			return nil, fmt.Errorf("createSolver: %w", err)
		}
		return pm.NewSolver(o), nil
	case "parallel":
		o := pl.NewOptions(*end, runtime.NumCPU())
		return pl.NewSolver(o), nil
	default:
		return nil, errors.New("createSolver: unknown options type provided")
	}
}
