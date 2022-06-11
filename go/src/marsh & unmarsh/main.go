// marsh & unmarsh project main.go
package main

import (
	"encoding/json"
	"fmt"
)

//Marshalling

type Emp struct {
	Name string
	Age  int
	Loc  string
}

func main() {

	emp_detials := Emp{Name: "mani", Age: 31, Loc: "Banglore"}
	emp, _ := json.Marshal(emp_detials)
	fmt.Println(string(emp))

	//unmarshall starts from here
	bytes := []byte(emp)
	var det details
	json.Unmarshal(bytes, &det)
	fmt.Println(det.Name)

}

//Unmarshalling

type details struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Loc  string `json:"loc"`
}

/*
package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Enter a positive integer : ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ { // assigning 1 to i
		sum += i // sum = sum + i
	}
	fmt.Print("Sum :", sum)
}


//SUM OF SQUARES (REMOTE WORK FROM SINGAPORE)
package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	for i := 1; i <= 3; i++ {
		//fmt.Println("squares:", i*i)
		sum = sum + (i * i)
		if i == 3 {
			fmt.Println("sum of squares:", i, sum)
		}
	}
}
*/
