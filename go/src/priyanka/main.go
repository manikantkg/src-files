// priyanka project main.go
package main

import (
	"fmt"
)

/*
func main() {
	tec := [5]int{1, 2, 3, 4, 5}

	for i, v := range tec {

		fmt.Println(i, v)
	}
	fmt.Println("Hello World!")
}


package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}


//function closures

func main() {

	// Declaring the variable
	GFG := 0

	// Assigning an anonymous
	// function to a variable
	counter := func() int {
		GFG += 1
		return GFG
	}

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

}


// newCounter function to
// isolate global variable
func newCounter() func() int {
    GFG := 0
    return func() int {
      GFG += 1
      return GFG
  }
}
func main() {
    // newCounter function is
    // assigned to a variable
    counter := newCounter()

    // invoke counter
    fmt.Println(counter())
    // invoke counter
    fmt.Println(counter())


}

*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
