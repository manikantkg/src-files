// emphasisi project main.go
/*
package main

import (
	"fmt"
	"sync"
)

func abc() {

	fmt.Println("HELLO WORLD")
	defer wg.Done()
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 2; i++ {
		go abc()
		wg.Add(1)
		wg.Wait()

	}
}

*/

package main

import (
	"fmt"
)

func abc() {
	var response string
	arr := response
	//	arr := "manikanta"
	fmt.Println(arr)
}
func main() {
	abc()
	fmt.Println("MYFRIEND")
}
