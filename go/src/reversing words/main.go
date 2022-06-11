// reversing words project main.go
// write a logic input = "My Name Is Sharath"  , output= "Sharath Is Name My"

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var a, b []byte = "m", "a"
	if bytes.Compare(a, b) < 0 {
		fmt.Println("a less b") //a less b
	}
}

/*
func reverseWords(s string) string {
	bytes := []byte(s)
	// remove additional spaces
	// pos - is a position for next found letter
	pos := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == 32 {
			continue
		}
		if pos > 0 && bytes[pos-1] != 32 && bytes[i-1] == 32 {
			bytes[pos] = 32
			pos++
		}
		bytes[pos] = bytes[i]
		pos++
	}
	bytes = bytes[:pos]

	//reverse whole string
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
	}

	//reverse words
	wordStartPos := 0
	for i := 0; i < len(bytes); i++ {
		// if this is the last letter in a word
		if i == len(bytes)-1 || bytes[i+1] == 32 {
			half := (i - wordStartPos + 1) / 2
			for pos := 0; pos < half; pos++ {
				bytes[wordStartPos+pos], bytes[i-pos] = bytes[i-pos], bytes[wordStartPos+pos]
			}
			// Since string was sanitized from additional spaces,
			// we could assume that next word will start in +2 symbols after last symbol of this word
			wordStartPos = i + 2
			i++
		}
	}
	return string(bytes)
}

func main() {
	re := reverseWords("My Name Is Sharath")
	fmt.Println(re)
}
*/
/*
func main() {
	// wrong approch
	var arr [9]int = [9]int{1, 3, 4, 3, 2, 3, 4, 0, 2}
	var item int
	for i := 0; i < len(arr); i++ {
		for j := 0; i < len(arr); j++ {
			var count int = 0
			if i == j {
				continue
			}
			if arr[i] == arr[j] {
				count++
				item = arr[j]
			}
			fmt.Println(count, item)
		}
		fmt.Println(arr)
	}

}


func main() {
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
}
*/
