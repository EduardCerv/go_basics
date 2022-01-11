package main

import "fmt"

func main() {
	//define and after assign
	var name  string
	name = "Eduardo"

	//define and assign at the same time
	var name_2 string = "Miry"

	//assign without define go automatically predict the kind of data
	age := 31

	fmt.Println(name, name_2, age)

	//values by default assigned by each kind of data
	var a int
	var b string
	var c int64
	var d bool
	var e float64

	fmt.Println(a, b, c, d, e)

	const pi = 3.1416

	fmt.Println(pi)
}