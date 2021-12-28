package main

import "fmt"

func main() {
	//definir y despues asignar
	var name  string
	name = "Eduardo"

	//definir y asignar a la vez
	var name_2 string = "Miry"

	//assignar sin definir automanticamente el compilador lo predice
	age := 31

	fmt.Println(name, name_2, age)

	// valores por defecto de cada tipo de variable
	var a int
	var b string
	var c int64
	var d bool
	var e float64

	fmt.Println(a, b, c, d, e)

	const pi = 3.1416

	fmt.Println(pi)
}