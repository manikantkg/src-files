// logicsown project main.go
package main

import (
	"fmt"
)

/*
func fizzBuzz(n int) {
	// Write your code here
	for n := 1; n <= 10; n++ {
		fmt.Print(n*3, " ")
		if (n%3 == 0) && (n%5 == 0) {
			fmt.Println("FizzBuzz")
		} else if n%3 == 0 {
			fmt.Println(" Fizz ")

		} else if n%5 == 0 {
			fmt.Println(" Buzz ")

		} else if (n/3 != 0) && (n/5 != 0) {
			fmt.Println("Nil")
		}
	}
}

func main() {
	fizzBuzz(15)
}


func count() {
	for i := 1; i < 10; i++ {
		fmt.Println("factorial", i*(i+1))
	}
}
func main() {
	count()
}
*/

type Vehicle interface {
	Structure() []string // Common Method
	Speed() string
}

type Car string

func (c Car) Structure() []string {
	var parts = []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	fmt.Println(parts)
	return parts
}

func (c Car) Speed() string {
	return "200 Km/Hrs"
}
func main() {

	Structure()
}
