// io library project main.go
/*
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "MANIKANTA  &   ANUSHA\n")
	}
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
*/
/*
package main

import (
	"fmt"
)

type car struct {
	gas_pedal      uint64
	breach_pedal   uint64
	steering_wheel int16
	top_speed_kmh  float64
}

func main() {
	a_car := car{
		gas_pedal:      22341,
		breach_pedal:   0,
		steering_wheel: 12561,
		top_speed_kmh:  225.0,
	}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal)
}
*/
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name string
	age  int
	loc  string
}

func main() {
	emp := person{name: "mani", age: 31, loc: "bang"}
	emp_det, _ := json.Marshal(emp)
	fmt.Println(string(emp_det))
}
