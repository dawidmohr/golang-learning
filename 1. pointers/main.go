package main

import "fmt"

func main() {
	// in Go, we don't using semicolons at end of line
	// u must also pass a variable type, like string in this example. God bless JS
	var nameOfJediKnight string = "Luke Skywalker"

	// variable must be used in this language. Otherwise program won't compile
	fmt.Print("I'm " + nameOfJediKnight + "and I'm Jedi Knight")

	// When u add "&" before variable name, u pass pointer to the variable
	joinTheDarkSide(&nameOfJediKnight)

	// Changed name by pointer in func
	fmt.Print("I'm " + nameOfJediKnight + "and I'm Jedi Knight")
}

// by *string we pass pointers with chosen type
func joinTheDarkSide(name *string) {
	// just allocating new string in memory by pointer
	*name = "Darth Wader"
}
