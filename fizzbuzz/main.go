package main

import (
	"flag"
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

	s := createSolver()

	if results, err := s.Solve(*end); err != nil {
		log.Fatal(err)
	} else {
		io.WriteString(os.Stdout, strings.Join(results, ", "))
	}
}

func createSolver() solver.FizzBuzzSolver {
	switch *approach {
	case "brute":
		return bf.NewSolver()
	case "parametrised":
		terms := pm.NewTermCollection(
			pm.NewTerm(func(v int) bool { return v%3 == 0 }, "Fizz", pm.LowPriority, false),
			pm.NewTerm(func(v int) bool { return v%5 == 0 }, "Buzz", pm.LowPriority, false),
			pm.NewTerm(func(v int) bool { return v%15 == 0 }, "FizzBuzz", pm.HighPriority, true),
		)

		o := pm.NewOptions(*begin, terms)
		return pm.NewSolver(o)
	case "parallel":
		o := pl.NewOptions(runtime.NumCPU())
		return pl.NewSolver(o)
	default:
		log.Fatal("createSolver: unknown options type provided")
		return nil
	}
}
