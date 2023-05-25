// Implements solution using worker pool
package parallel

import (
	"strconv"
	"sync"
)

type Options struct {
	// Number of workers
	wc int
}

func NewOptions(w int) *Options {
	return &Options{wc: w}
}

type Solver struct {
	o *Options
}

func NewSolver(opt *Options) *Solver {
	return &Solver{o: opt}
}

// Solve takes n as input, to generate values till it reached starting from 1
// Refer to https://leetcode.com/problems/fizz-buzz/description/
// Returns no errors
func (s *Solver) Solve(n int) ([]string, error) {
	var res = make([]string, 0, n)

	// Because we know the size of the collection, we can safely use buffered channels
	jobs := make(chan int, n)
	results := make(chan string, n)

	// Create worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < s.o.wc; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// Send jobs to the worker goroutines
	go func() {
		for i := 1; i <= n; i++ {
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
