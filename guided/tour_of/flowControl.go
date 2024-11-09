package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	// defer statement does not execute till the surrounding function returns
	defer fmt.Println("deferred till the end")

	// only looping construct but bit strange without bracket grouping the loop
	// pretty much the same
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("one", sum)

	// init and post statements are optional
	// the loop limiter is required
	// essentially a while loop
	sum1 := 1
	for sum1 < 10 {
		sum1 += sum1
	}
	fmt.Println("two", sum1)

	// infinite loop
	// for {}

	// if is the same doesnt need () but require {}
	// if x > 10 {
	//}
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow_elif(3, 2, 10))
	fmt.Println(sqr(2, 20))

	// standard switch operation
	// evaluates cases from top to bottom
	// switch can run with out condition
	switch_cases()
	// switch {} used to run lond elif chain
	switch_elif()

	// deferred function calls are pushed into a stack
	// executed LIFO order
	// https://go.dev/blog/defer-panic-and-recover
	defer fmt.Println("deferr stack LIFO")
}

// if can start with a short statement to execute before the condition
// var declared here are only available in scope
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// vars declared inside if short statement are also available inside else
func pow_elif(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println("%g >= %g\n", v, lim)
	}
	return lim
}

// ex: square root using loops
func sqr(x float64, p int) float64 {
	z := 1.0
	for c := 1; c < p; c++ {
		z -= (z*z - 2) / (2 * x)
	}
	return z
}

func switch_cases() {
	fmt.Printf("\nGo runs on - ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("%s . \n", os)
	}
}

func switch_elif() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
