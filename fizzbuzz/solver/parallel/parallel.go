package parallel

import (
	"strconv"
	"sync"
)

type Options struct {
	// Specifies the end of execution, when incrementing loop index reaches this value
	N int
	// Number of workers
	WorkersCount int
}

func NewOptions(n int, w int) *Options {
	return &Options{N: n, WorkersCount: w}
}

type Solver struct {
	o *Options
}

func NewSolver(opt *Options) *Solver {
	return &Solver{o: opt}
}

func (s *Solver) Solve() ([]string, error) {
	var res = make([]string, 0, s.o.N)

	// Because we know the size of the collection, we can safely use buffered channels
	jobs := make(chan int, s.o.N)
	results := make(chan string, s.o.N)

	// Create worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < s.o.WorkersCount; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// Send jobs to the worker goroutines
	go func() {
		for i := 1; i <= s.o.N; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Collect results from the worker goroutines
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	for r := range results {
		res = append(res, r)
	}

	return res, nil
}

func worker(jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range jobs {
		switch {
		case num%15 == 0:
			results <- "FizzBuzz"
		case num%3 == 0:
			results <- "Fizz"
		case num%5 == 0:
			results <- "Buzz"
		default:
			results <- strconv.Itoa(num)
		}
	}
}
