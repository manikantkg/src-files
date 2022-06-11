// setget project main.go
package main

import (
	"fmt"
)

type Person struct {
	name string

	phoneno int
}

func (p *Person) setmethod(na string) {

	p.name = na

}

func (p Person) getmethod() string {

	return p.name
}
func main() {
	var sharath Person
	sharath.setmethod("Go Lang")

	//fmt.Println(sharath)
	fmt.Print(sharath.getmethod())
}
