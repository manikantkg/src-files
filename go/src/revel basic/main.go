// revel basic project main.go
package main

import (
	"fmt"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func main() {
	fmt.Println("manikanra")
}
func (c App) Index() revel.Result {
	fmt.Println("Hello world!")
	return c.Render()
}
