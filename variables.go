package main

import "fmt"

func main() {
	var a = "initial" // simple declaration and initializing of a single variable with var
	fmt.Println(a)

	var b, c int = 1, 2 // declaring and initializing multiple variables at once
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // simple declararion of variables without initializing it, the come wuth zero values
	fmt.Println(e)

	f := "apple" // := is shorthand method of declaring and initializing variable
	fmt.Println(f)
}
