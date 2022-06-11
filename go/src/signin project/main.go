// signin project project main.go

// sign desk project main.go
package main

import (
	"fmt"
)

/*
type details struct {
	name string

	roll int64
}

func (p *details) setinfo() {
	fmt.Println("enter details")
	fmt.Scan(&p.name, &p.roll)

}

func (p *details) getinfo() {

	fmt.Println(p.name, p.roll)

}

func main() {
	student := details{}

	student.setinfo()
	student.getinfo()
}
*/
func student1(m chan int) {
	fmt.Println("hi swathi")

	m <- 123
}
func student2(m chan int) {
	fmt.Println("helo sharath", <-m)

}

func main() {

	m := make(chan int)

	go student1(m)

	go student2(m)

}
