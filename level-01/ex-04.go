package main

import "fmt"

type myType int

var x myType

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
}
