// logic project main.go
package main

import (
	"fmt"
)

/*
func main() {
	//**********febbinoci series**********
	var num int = 10
	var n3, n1, n2 int = 0, 0, 1
	for i := 1; i <= num; i++ {
		fmt.Println(n1)
		n3 = n1 + n2
		n1 = n2
		n2 = n3

	}
}


const (
	i = 7
	j
	k
)
*/ /*
func main() {
	//fmt.Println(i, j, k)


			//**********PASCAL
				var rows int
				var temp int = 1
				fmt.Print("Enter number of rows : ")
				fmt.Scan(&rows)

				for i := 0; i < 5; i++ {

					for j := 1; j <= rows-i; j++ {
						fmt.Print(" ")
					}

					for k := 0; k <= i; k++ {

						if k == 0 || i == 0 {
							temp = 1
						} else {
							temp = temp * (i - k + 1) / k
						}

						fmt.Printf(" %d", temp)
					}
					fmt.Println("")

				}


		//***********PYRAMID
		var rows int = 5
		var k int

		for i := 1; i <= rows; i++ {
			k = 0
			for space := 1; space <= rows-i; space++ {
				fmt.Print("  ")
			}
			for {
				fmt.Print("* ")
				k++
				if k == 2*i-1 {
					break
				}
			}
			fmt.Println("")
		}
	}
	//********TRIANGLE NUMBERS
	var rows = 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
//***********FIBBINOCI SERIES
func main() {
	var n int
	t1 := 0
	t2 := 1
	nextTerm := 0

	fmt.Print("Enter the number of terms : ")
	fmt.Scan(&n)
	fmt.Print("Fibonacci Series :")
	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(" ", t1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", t2)
			continue
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		fmt.Print(" ", nextTerm)
	}
}


//********SUM OF NATURAL NUMBERS

func main() {
	var n, sum int
	fmt.Print("Enter a positive integer : ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ { // assigning 1 to i
		sum += i // sum = sum + i
	}
	fmt.Print("Sum : ", sum)
}
*/

func main() {
	var number, remainder, temp int
	var reverse int = 0

	fmt.Print("Enter any positive integer : ")
	fmt.Scan(&number)

	temp = number

	// For Loop used in format of While Loop
	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break // Break Statement used to exit from loop
		}
	}

	if temp == reverse {
		fmt.Printf("%d is a Palindrome", temp)
	} else {
		fmt.Printf("%d is not a Palindrome", temp)
	}

}
