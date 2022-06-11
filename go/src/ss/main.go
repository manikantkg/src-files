// ss project main.go

package main

import (
	"fmt"
)

/*
func square() {
	var l int
	a := map[int]interface{}{

		4: 6,
	}
	area2 := (l * l)
	fmt.Println("area:", area2)
}
func rect() {
	var l, b int

	area1 := (l * b)
	fmt.Println("area:", area1)
}

func main() {
	square()
	rect()
}
*/
/*
//practice

func arr(a []string) {
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i], "hello")
	}
	fmt.Println("******")
	fmt.Println(len(a), cap(a), a)
	//a = append(a, "m", "a", "n", "i")
}
func main() {
	arr([]string{"m", "a", "n", "i"})
}
*/

//ACCESSING OF A STRUCTURE

type char struct {
	name string
	age  int
}

func (p *char) details() {
	p.name = " MANI"
	p.age = 31
}
func main() {

	var mani char = char{"mani", 31}

	var manu char = char{"manu", 21}

	var sree char = char{"sree", 65}

	fmt.Println(mani, manu, sree)
}
