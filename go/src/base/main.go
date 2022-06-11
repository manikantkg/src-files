// base project main.go
package main

import (
	"fmt"
)

func main() {

	var a, b int = 2, 3
	var c = []string{"hello", "how", "are", "you"}
	fmt.Println(a, b, "\n", c)

	fmt.Println("MANIKANTA")

	if 7%2 == 0 {
		fmt.Println("7 is EVEN")
	} else {
		fmt.Println("7 is ODD")
	}

	if num := 10; num <= 10 {
		fmt.Println("num is num")
	} else if num < 10 {
		fmt.Println("num is odd")
	} else {
		fmt.Println("num is has only one digit")
	}

	const e = 2.33
	fmt.Println(e)
	const p, q = 345, 675
	fmt.Println(p, "\n", q)
}

//8897211372- gangadhar tailor
