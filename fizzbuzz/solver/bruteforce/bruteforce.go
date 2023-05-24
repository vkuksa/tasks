package bruteforce

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

const (
	fb = "FizzBuzz"
	f  = "Fizz"
	b  = "Buzz"
)

type BruteForce struct {
}

func NewSolver() *BruteForce {
	return &BruteForce{}
}

func (bf *BruteForce) Solve(w io.Writer, num int) error {
	if w == nil {
		return errors.New("brutforce solve: passing nil writer")
	}

	var err error
	var str string
	var n int

	for i := 1; i <= num; i++ {
		switch {
		// Deliberately omitting fallthrough here to increase readability
		case i%15 == 0:
			str = fb
		case i%3 == 0:
			str = f
		case i%5 == 0:
			str = b
		default:
			str = strconv.Itoa(i)
		}

		n, err = io.WriteString(w, str)
		if err != nil {
			return fmt.Errorf("brutforce solve: %w", err)
		}
		if n != len(str) {
			return fmt.Errorf("brutforce solve: %s was not properly written", str)
		}
	}

	return nil
}
