package main

// import packages with path
// convention - package name is the same as the last element of the import path
// can either group import like this or seperately
// group is called factored
import (
	"fmt"
	"math"
	"math/rand"
)

// variables, last key is type
// can be package(global) or function level
var c, py, java bool

// entry point, this is called a package??
func main() {
	fmt.Println("Hello random number", rand.Intn(10))
	fmt.Println("Hello c++ string interpolation %g number", math.Sqrt(7))

	// a name is exported if it begins with a capital letter
	fmt.Println("exported name", math.Pi)

	fmt.Println("casual %g + %g = %g", 1, 2, add(1, 2))

	multiA, multiB := swap("hello", "world")
	fmt.Println(multiA, multiB)

	fmt.Println(split(17))

	// function level variables
	var i int
	fmt.Println(i, c, py, java)

	// var can also chain initializers
	var cpp, chash = true, "egg"
	fmt.Println(cpp, chash)

	// short hand var declaration only in a func
	// := also infered type, cannot beused outside a func
	// on package level, every statement begins with a keyword
	k := 3            // int
	m := 3.14         // float64
	n := 0.123 + 0.5i // complex128
	fmt.Println(k, m, n)

	// zero values
	// var declared without initial value are given zero value
	// 0 numeric
	// false for boolean
	// "" for empty string

	// T(v) to convert value of v into type T

	// cannot be declared using :=
	const egg = "egg"
	fmt.Println(egg)
}

// functions seems self expolanatory
// more on types https://go.dev/blog/declaration-syntax
// when 2 or more consecutive named params share a type
// you can omit the type exept the last
// func add(x, y int) int {
func add(x int, y int) int {
	return x + y
}

// can return multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// return can return named values, theyre treated as named vars named at the top
// naked return - returns the named args in the named return value
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/*
	Every go program made up of packages
	Program start running in package main
*/

/* basic types
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/
