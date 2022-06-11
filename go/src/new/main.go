// new project main.go

package main

import (
	"fmt"
)

/*
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	http.ListenAndServe(":8080", router)
}


func main() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	fmt.Println(s)
}


func main() {


}


package main

import (
	"fmt"
	"time"
)

func printnumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printletters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go printnumbers()
	go printletters()
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Printing from main")
}
*/
func main() {
	var v *int
	v = new(int)
	*v = 8
	fmt.Println(*v)
	//mv := make(map[string]string)
	//fmt.Println("mv: %p %#v \n", &mv, mv)

	var mv *map[string]string
	fmt.Printf("mv: %p %#v \n", &mv, mv) //mv

	const (
		C0 = iota
		C1
		C2
		C3
	)
	fmt.Println(C0, C1, C2, C3)

	/*
			const (
		    North Direction = iota
		    East
		    South
		    West
		)

		func (d Direction) String() string {
		    return [...]string{"North", "East", "South", "West"}[d]
		}
	*/
	var a = []int{1, 2, 3, 4}

	fmt.Println(a{1: 1})
}
