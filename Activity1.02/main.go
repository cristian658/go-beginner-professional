package main

import (
	"fmt"
)

var foo string = "bar"

func main() {
	fmt.Println("Hola mundo")

	var baz string = "qux"

	fmt.Println(foo, baz)
}
