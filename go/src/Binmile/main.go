// Binmile project main.go
/*
package main

import (
	"fmt"
	"strings"
)

func occur(str string) map[string]int {
	word := strings.Fields(str)
	count := make(map[string]int)

	//

	for _, choose := range word {
		_, ok := count[choose]
		if ok {
			count[choose] += 1
		} else {
			count[choose] = 1
		}

	}
	return count
}

func main() {
	str := "h e l l o w o r l d"
	for index, element := range occur(str) {
		fmt.Println(index, " =>", element)
	}
}


package main

import (
	"fmt"
)

func print(arr []int) int {

	element := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if element[arr[i]] == true {

			return arr[i]

		} else {
			element[arr[i]] = true
		}
	}
	return -1
}
func main() {

	fmt.Println(print([]int{5, 2, 3, 3, 4, 1, 2}))
}
*/
/*
package main

import "fmt"

func duplicateInArray(arr []int) int {
	visited := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return arr[i]
		} else {
			visited[arr[i]] = true
		}
	}
	return -1
}

func main() {
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 2}))
	fmt.Println(duplicateInArray([]int{1, 4, 4, 2, 3}))
}


*/

package main

import "fmt"

func main() {
	// local variable definition
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	//* calling a function to swap the values.
	//* &a indicates pointer to a ie. address of variable a and
	//* &b indicates pointer to b ie. address of variable b.
	//*
	swap(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)

	swapp()
}
func swap(x *int, y *int) {
	var temp int
	temp = *x // save the value at address x
	*x = *y   //* put y into x
	*y = temp //* put temp into y
}

//swapp
func swapp() {
	a := 10
	b := 20
	var temp int

	temp = a
	a = b
	b = temp
	fmt.Println(a, b)

}
