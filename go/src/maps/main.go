// maps project main.go
package main

import (
	"fmt"
)

/*
func main() {

	m := make(map[string]int)

	m["mani"] = 345
	m["anu"] = 654
	m["sree"] = 65
	m["lakshmi"] = 60
	fmt.Println(m)

	_, prs := m["lakshmi"]
	fmt.Println(prs)
	delete(m, "lakshmi")
	fmt.Println(m)

}
*/

// DUPLICATE NUMBERS COUNTING

func main() {
	var arr []int = []int{1, 3, 4, 3, 2, 4, 4, 0, 2, 4, 3}

	mm := make(map[int]int)

	for _, v := range arr {

		// fmt.Println(i, v)

		mm[v] = mm[v] + 1
		// mm[1] = mm[1]+1                mm[1] = 0 + 1

		// mm[3] = mm[3] + 1              mm[3] = 0 + 1

		// mm[4] = mm[4] + 1              mm[4] = 0 + 1

		// mm[3] = mm[3] + 1              mm[3] = 1+ 1

	}

	fmt.Println(mm)

}

/*

func main() {

	var arr = []string{"hello", "java", "hello", "golang", "hello", "golang"}

	var mapp = make(map[string]int)
	for _, v := range arr {

		mapp[v] = mapp[v] + 1

	}
	fmt.Println(mapp)
}
*/
