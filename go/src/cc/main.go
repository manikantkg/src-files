// cc project main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	
	
)


//name age add // object 
type Employee struct{
	Name string `json:"Name"`
	Age int `json:"Age"`
	Address string `json:"Address"`
	
}
var *db.gorm

func main() {
	
	
	
	var Emp = []Employee
	Employee{Name:"rahul", Age:30, Address:"HYD"},
	r:= gin.Default()
	r.POST (func (c *gin.Context){
		c.Json(c.StatusOK,200)
	})
	 c.Json(Employee(&Emp))
}
