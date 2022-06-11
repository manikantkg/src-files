// array class work project main.go
package main

import (
	"fmt"
	//"strings"
	//"github.com/revel/revel/session"
)

/*
func array() {

	var studentname [5]string = [5]string{"venkat", "manikanta", "basava", "sree", "manu"}

	var friendlist [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("slice lene", len(studentname)) // Length of an array

	fmt.Println(studentname[0], studentname[2])

	fmt.Println(friendlist[0], friendlist[1])

	for i, v := range studentname {

		fmt.Println(i, v) // INDEX OF ARRAY
	}
}
func arr() {

	var arr [3]int = [3]int{1, 2, 3}

	var food [3]string = [3]string{"dosa", "chapati", "idly"}

	fmt.Println(food[0], food[2])

	for i := 0; i < 3; i++ {

		fmt.Println(arr[i])
	}

	var crr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(crr)

}
func practice() {

	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	var food [3]string = [3]string{"dosa", "idly", "chapati"}

	var friends [5]string = [5]string{"xx", "yy", "zz", "kk", "mm"}

	for i := 0; i < 5; i++ {
		fmt.Println(friends[i])

		//fmt.Println(i, v:= 0)
	}
	fmt.Println(len(arr), arr[0], arr[1])

	fmt.Println(food[0], food[2])

}

func pointers() {

	var code int = 5

	var p *int = &code

	fmt.Println("code block", &code, code)

	fmt.Println("code block", &p, p)

	var chocolate string = "diarymilk"

	var favouritem *string = &chocolate

	fmt.Println("chocolate block", &chocolate, chocolate)

	fmt.Println("favouritem block", &favouritem, favouritem)
}

func New() {

	var principal *string = new(string)

	*principal = "Mr.Venkat"

	var car *string = new(string)

	*car = "celerio"

	var jvt *[3]int = new([3]int)

	*jvt = [3]int{555, 666, 777}

	var teachers []string = make([]string, 3)

	teachers[0] = "sir"

	teachers[1] = "madam"

	fmt.Println(principal, &principal, *principal)

	fmt.Println(car, &car, *car)

	fmt.Println(jvt)

	var jvt1 *[3]string = new([3]string)

	*jvt1 = [3]string{"*c", "*cpp", "*java"}

	jvt1[0] = "*c"

	fmt.Println(*principal, teachers, jvt1, jvt)

	teachers = append(teachers, "java")
}
*/
func passlicearray(tec []string, j *int) {

	var jj[2] = [2]int{2, 33}

	for i, v := range tec {

		if v == "java" {

			fmt.Println(i, v)
		}

		fmt.Println(tec) // arr[]int = []int{1,2,3,4}
	}

}

func main() {
	passlicearray([]string{"java", &jj}) //, 22
	//array()
	//arr()
	//practice()
	//pointers()
	//New()
	//slice()
	//basic()
	//trials()

}

/*
func trials() {

	type Person struct {
		name string
		age  int
	}

	// our Team struct
	type Team struct {
		name    string
		players [2]Person
	}

	// declaring an empty 'Team'
	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{Person{name: "Forrest", age: 25}, Person{name: "Gordon"}}
	// declaring a team with players
	celtic := Team{name: "Celtic FC", players: players}
	verna := Team{name: "google", players: players}
	fmt.Println(celtic, verna)

}

func explore() {
	var i interface{} = "Manikanta"
	fmt.Printf("%v\n", i)

	fmt.Println("Hello World")

	var myint int8
	for i := 0; i < 129; i++ {
		myint += 1
	}
	fmt.Println(myint)

	var isTrue bool = true
	var isFalse bool = false
	// AND
	if isTrue && isFalse {
		fmt.Println("Both Conditions need to be True")
	}
	// OR - not exclusive
	if isTrue || isFalse {
		fmt.Println("Only one condition needs to be True")
	}
}
func basic() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //  index from 1 to 4
	fmt.Println(s)

	var a = 45

	// Initialization of pointer s with
	// memory address of variable a
	var t *int = &a

	fmt.Println(t, *t, a)

	var y = 458

	// taking a pointer variable using
	// var keyword without specifying
	// the type
	var p = &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)

	// this is dereferencing a pointer
	// using * operator before a pointer
	// variable to access the value stored
	// at the variable at which it is pointing
	fmt.Println("Value stored in y,(*p) = ", *p)

}


func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

/*
syntax a pointers()

*int/  * string    (*datatype)'*' operator indicates  how many blocks of memory it OCCUPIES   {DEREERENCING }

'&' OPERATOR INDICATES TO ADDRESS OF THE PARTICULAR BLOCK


func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

*/

/*
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
*/
/*
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}*/

/*
//DIFFERENCE BETWEEN LENGTH & CAPACITY

package main

import "fmt"

func main() {
   // create a slice
   numbers := []int{0,1,2,3,4,5,6,7,8}
   printSlice(numbers)

   // print the original slice
   fmt.Println("numbers ==", numbers)

   // print the sub slice starting from index 1(included) to index 4(excluded)
   fmt.Println("numbers[1:4] ==", numbers[1:4])

   // missing lower bound implies 0
   fmt.Println("numbers[:3] ==", numbers[:3])

   // missing upper bound implies len(s)
   fmt.Println("numbers[4:] ==", numbers[4:])

   numbers1 := make([]int,0,5)
   printSlice(numbers1)

   // print the sub slice starting from index 0(included) to index 2(excluded)
   number2 := numbers[:2]
   printSlice(number2)

   // print the sub slice starting from index 2(included) to index 5(excluded)
   number3 := numbers[2:5]
   printSlice(number3)

}
func printSlice(x []int){
   fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x),x)
}

*/
