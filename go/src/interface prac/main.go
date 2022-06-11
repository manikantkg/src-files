// interface prac project main.go

/*
package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	fmt.Println(a.String())
}

*/
//STRING FORMATTING

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
	//If the value is a struct, the %+v variant will include the struct’s field names.

	fmt.Printf("struct2: %+v\n", p)
	//The %#v variant prints a Go syntax representation of the value,
	//i.e. the source code snippet that would produce that value.

	fmt.Printf("struct3: %#v\n", p)
	//To print the type of a value, use %T.

	fmt.Printf("type: %T\n", p)

	fmt.Printf("bool: %t\n", true)

	fmt.Printf("int: %d\n", 123)

	fmt.Printf("bin: %b\n", 14)

	fmt.Printf("char: %c\n", 33)
	//%x provides hex encoding.

	fmt.Printf("hex: %x\n", 456)

	fmt.Printf("float1: %f\n", 78.9)

	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"")

	fmt.Printf("str2: %q\n", "\"string\"")

	fmt.Printf("str3: %x\n", "hex this")

	fmt.Printf("pointer: %p\n", &p)

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	//You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	//To left-justify use the - flag as with numbers.

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
	//So far we’ve seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere.

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	//You can format+print to io.Writers other than os.Stdout using Fprintf.

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
