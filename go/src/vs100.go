package main

import (
	"fmt"
)

type A struct {
	name string
	age  int
	loc  string
}

func (BB *A) PERSON() {
	BB.name = "mani"
	BB.age = 30
	BB.loc = "Bang"

	fmt.Println(BB.name, BB.age, BB.loc)
}
func main() {
	var mani A
	mani.PERSON()
	var sai A = A{"sai", 25, "hyd"}
	SHARAT := A{"sharat", 28, "HYD"}
	fmt.Println("manikanta", SHARAT, sai)
}
