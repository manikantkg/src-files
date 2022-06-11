// nitya project main.go
package main

import (
	"fmt"
)

func main() {
	s := make(map[string]struct{})

	var exists = struct{}{}

	s["James"] = exists
	s["David"] = exists
	s["Peter"] = exists

	/*
		if _, ok := s["Peter"]; ok {
			fmt.Println("Peter is in the set.") // This will be printed
		}

		if _, ok := s["George"]; ok {
			fmt.Println("George is in the set.") // This won't be printed
		}
	*/
	delete(s, "Peter")

	if _, ok := s["Peter"]; ok {
		fmt.Println("Peter is in the set") // Now this won't be printed
	}
	for v := range s {
		fmt.Println(v)
	}
}
