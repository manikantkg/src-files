// slice project main.go
package main

import (
	"fmt"
)

func main() {

	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("slice", s)

	s = append(s, "subbu")
	fmt.Println("len:", len(s), "cap:", cap(s))
	fmt.Println(s)

	c := make([]string, 4)
	copy(c, s)
	fmt.Println(c)

	d := s[2:4]
	fmt.Println(d)

	e := s[1:3]
	fmt.Println(e)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
