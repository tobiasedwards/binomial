package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		printUsage()
	} else {
		n, p, x := getArguments()
		prob, err := cumulativeBinomial(n, p, x)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%.6f\n", prob)
	}
}

// An Interval represents an interval from one integer to another.
// It can be open or closed at either end.
// The notation for an Interval that only contains one number, n, is [n, n+1).
type Interval struct {
	openLeft, openRight bool
	left, right         int64
}

// Start returns the start point when looping through i..
// If i is openLeft, left + 1 is returned but if i is not openLeft, left is
// returned.
func (i *Interval) Start() int64 {
	if i.openLeft {
		return i.left + 1
	} else {
		return i.left
	}
}

// End returns the end point when looping through i.
// If i is openRight, right - 1 is returned but if i is not openRight, right it
// returned.
func (i *Interval) End() int64 {
	if i.openRight {
		return i.right - 1
	} else {
		return i.right
	}
}

// printUsage prints the usage of this command line tool to the terminal.
func printUsage() {
	usage := "Usage: binomial n p x\n" +
		"  n: Number of trials. Integer greater than 0\n" +
		"  p: Probability of success. Float between 0 and 1\n" +
		"  x: Number of successes required.\n" +
		"     Integer between 0 and n of an open or closed interval using the\n" +
		"     following notation - [a,b], (a,b) or a combination of the two.\n"

	fmt.Print(usage)
}

// getArguments gets the command line arguments n, p and x.
// If there is an error during this process getArguments will log it and exit.
func getArguments() (n int64, p float64, x Interval) {
	n, err := validateN()
	if err != nil {
		log.Fatal(err)
	}

	p, err = validateP()
	if err != nil {
		log.Fatal(err)
	}

	x, err = validateX()
	if err != nil {
		log.Fatal(err)
	}

	return
}

// validateN returns p from command line arguments and validates it.
func validateN() (n int64, err error) {
	if len(os.Args) < 2 {
		err = errors.New("positional argument n not provided")
		return
	}

	n_string := os.Args[1]
	n, err = strconv.ParseInt(n_string, 10, 64)
	if err != nil {
		return
	}

	if n <= 0 {
		err = fmt.Errorf("n must be greater than 0; %d provided", n)
		return
	}

	return
}

// validateP returns p from command line arguments and validates it.
func validateP() (p float64, err error) {
	if len(os.Args) < 3 {
		err = errors.New("positional argument p not provided")
		return
	}

	p_string := os.Args[2]
	p, err = strconv.ParseFloat(p_string, 64)
	if err != nil {
		return
	}

	if p > 1 || p < 0 {
		err = fmt.Errorf("p must be in interval [0,1]; %f provided", p)
		return
	}

	return
}

func validateX() (x Interval, err error) {
	if len(os.Args) < 4 {
		err = errors.New("positional argument x not provided")
		return
	}

	x_string := os.Args[3]
	x_64, err := strconv.ParseInt(x_string, 10, 64)

	// If there is an error, then not just an integer, expect interval
	if err != nil {
		x, err = parseInterval(x_string)
		return
	} else {
		// Return the interval [x, x+1), which will just perform x
		x = Interval{false, true, x_64, x_64 + 1}
		return
	}
}

// parseInterval returns an interval from a given interval_string.
// If the string does not match notation correctly an error will be returned.
//
// Some examples of correct notation:
//
//   [2,3]
//   (23, 45]
//   (3,3)
//
// The interval_string must:
//
//   start with either '[' or '('
//   end with ']' or ')'
//   contain integers only
//   have left boundary <= right boundary
//
func parseInterval(interval_string string) (interval Interval, err error) {
	if strings.HasPrefix(interval_string, "[") {
		interval.openLeft = false
	} else if strings.HasPrefix(interval_string, "(") {
		interval.openLeft = true
	} else {
		err = errors.New("interval string must begin with '[' or '('")
		return
	}

	if strings.HasSuffix(interval_string, "]") {
		interval.openRight = false
	} else if strings.HasSuffix(interval_string, ")") {
		interval.openRight = true
	} else {
		err = errors.New("interval string must end with ']' or ')'")
		return
	}

	interval_string = strings.Trim(interval_string, "[]()")
	boundary_strings := strings.Split(interval_string, ",")

	if len(boundary_strings) != 2 {
		err = fmt.Errorf("2 numbers numbers required; %d given",
			len(boundary_strings))
		return
	}

	left_string := boundary_strings[0]
	left_string = strings.Trim(left_string, " ")
	left, err := strconv.ParseInt(left_string, 10, 64)
	if err != nil {
		return
	}

	interval.left = left

	right_string := boundary_strings[1]
	right_string = strings.Trim(right_string, " ")
	right, err := strconv.ParseInt(right_string, 10, 64)
	if err != nil {
		return
	}

	interval.right = right

	if interval.left > interval.right {
		err = errors.New("left boundary greater than right")
		return
	}

	return
}

// factorial recusively calculates and returns the factorial of n.
// An error is returned if n is less than 0.
func factorial(n int64) (int64, error) {
	if n < 0 {
		return -1, fmt.Errorf("n cannot be less than 0; %d provided", n)
	}

	if n == 0 || n == 1 {
		return n, nil
	}

	// m represents factorial of n-1
	m, _ := factorial(n - 1)
	return n * m, nil
}

// combination returns the number of possible combinations of r items in a set
// of n items when order is not important.
// An error is returned if n or r are less than 0.
func combination(n, r int64) (int64, error) {
	if n < 0 {
		return -1, fmt.Errorf("n cannot be less than 0; %d provided", n)
	} else if r < 0 {
		return -1, fmt.Errorf("r cannot be less than 0; %d provided", r)
	}

	if r == 0 || r == n {
		return 1, nil
	} else if r > n {
		return 0, nil
	}

	n_fac, _ := factorial(n)
	r_fac, _ := factorial(r)
	n_less_r_fac, _ := factorial(n - r)
	return n_fac / (r_fac * n_less_r_fac), nil
}

// binomial calculates the probability of a single case X = x with parameters n
// and p.
func binomial(n, x int64, p float64) (float64, error) {
	combination, err := combination(n, x)
	if err != nil {
		return -1.0, err
	}
	return float64(combination) * math.Pow(p, float64(x)) *
		math.Pow(1-p, float64(n-x)), nil
}

// cumulativeBinomial calculated the probability of x in interval i with
// parameters n and p.
func cumulativeBinomial(n int64, p float64, i Interval) (float64, error) {
	var probability float64
	for x := i.Start(); x <= i.End(); x++ {
		b, err := binomial(n, x, p)
		if err != nil {
			return -1.0, err
		}

		probability += b
	}

	return probability, nil
}
