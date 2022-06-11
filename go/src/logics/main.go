// logics project main.go
package main

import (
	"fmt"
	//"reflect"
	//"strings"
)

func main() {

	/*	var str string = "manu"
									for i, v := range str {
										fmt.Println(i, v)
									}
									fmt.Println("Hello World!")

									//**********SWAPPING*********

									var arr []int = []int{'m', 'a', 'n', 'u'}

									var arr2 []string = []string{"m", "a", "n", "u"}

									temp := arr[0]
									arr[0] = arr[2]
									arr[2] = temp

									temp1 := arr2[0]
									arr2[0] = arr2[1]
									arr2[1] = arr2[2]
									arr2[2] = arr2[3]
									arr2[3] = temp1

									fmt.Println("manikanta")
									fmt.Println(arr, arr2)

									//****************FACTORIAL****************

									fact := 1

									for i := 1; i <= 5; i++ {
										fact = fact * i
									}
									fmt.Println(fact)

										var st strings.Builder

										for i := 0; i < 10; i++ {
											st.WriteString("a ")
										}

										fmt.Println(st.String())

											x := [...]int{10, 20, 30}

											fmt.Println(reflect.ValueOf(x).Kind())
											fmt.Println(len(x))

								x := [5]int{1: 10, 3: 30}
								fmt.Println(x)

							intArray := [5]int{10, 20, 30, 40, 50}

							fmt.Println("\n---------------Example 1--------------------\n")
							for i := 0; i < len(intArray); i++ {
								fmt.Println(intArray[i])
							}

							fmt.Println("\n---------------Example 2--------------------\n")
							for index, element := range intArray {
								fmt.Println(index, "=>", element)

							}

							fmt.Println("\n---------------Example 3--------------------\n")
							for _, value := range intArray {
								fmt.Println(value)
							}

							j := 0
							fmt.Println("\n---------------Example 4--------------------\n")
							for range intArray {
								fmt.Println(intArray[j])
								j++
							}


						countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

						fmt.Printf("Countries: %v\n", countries)

						fmt.Printf(":2 %v\n", countries[:2])

						fmt.Printf("1:3 %v\n", countries[1:3])

						fmt.Printf("2: %v\n", countries[2:])

						fmt.Printf("2:5 %v\n", countries[2:5])

						fmt.Printf("0:3 %v\n", countries[0:3])

						fmt.Printf("Last element: %v\n", countries[len(countries)-1])

						fmt.Printf("All elements: %v\n", countries[0:len(countries)])
						fmt.Println(countries[:])
						fmt.Println(countries[0:])
						fmt.Println(countries[0:len(countries)])

						fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])



					//***********Length & Capacity of Slice***********

					var intSlice = []int{10, 20, 30, 40}
					var strSlice = []string{"India", "Canada", "Japan"}

					fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
					fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))

					var intSlic = []int{10, 20, 30, 40}

					fmt.Println(intSlic[0])
					fmt.Println(intSlic[1])
					fmt.Println(intSlic[0:4])
					fmt.Println(intSlic[1:3])

				//Change item value

				var strSlice = []string{"India", "Canada", "Japan"}
				fmt.Println(strSlice)

				strSlice[2] = "Germany"
				fmt.Println(strSlice)

			//****Removing item value*******

			var str = []string{"India", "Canada", "Japan", "Germany", "Italy"}
			fmt.Println(str)

			str = RemoveIndex(str, 3)
			fmt.Println(str)

			func RemoveIndex(s []string, index int) []string {
			return append(s[:index], s[index+1:]...)
			}


		//********Copy a slice*********
		a := []int{5, 6, 7} // Create a smaller slice
		fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))

		b := make([]int, 5, 10) // Create a bigger slice
		copy(b, a)              // Copy function
		fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

		fmt.Println("Slice B after copying:", b)
		b[3] = 8
		b[4] = 9
		fmt.Println("Slice B after adding elements:", b)
	*/

	//*********Tricks by using indexes**********

	var countries = []string{"india", "japan", "canada", "australia", "russia"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[4])
	fmt.Printf("Last element: %v\n", countries[len(countries)-1])
	fmt.Printf("Last element: %v\n", countries[4:])

	fmt.Printf("All elements: %v\n", countries[0:len(countries)])

	fmt.Printf("Last two elements: %v\n", countries[3:len(countries)])
	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])

	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])
}
