// array project main.go
package main

import (
	"fmt"
)

func main() {
	var arr [4]int
	fmt.Println(arr)

	arr = [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	arr[3] = 100
	fmt.Println(arr)

	fmt.Println("set:", arr)
	fmt.Println("get:", arr[3])

	var bb = []string{"hello", "how", "are", "you"}
	fmt.Println(bb)
	fmt.Println(bb[3])

	fmt.Println(len(arr), cap(arr))
	fmt.Println(len(bb), cap(bb))

	q := append(bb, "manikanta")
	fmt.Println(q)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// slices
	var mm = [10]int{1, 2, 3, 4}
	fmt.Println(mm)

	fmt.Println(len(mm), cap(mm))
	mm[7] = 100
	fmt.Println(mm)

}
