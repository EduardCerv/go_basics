package main

import "fmt"

func main() {
	a := 20
	b := 10

	result := a + b

	fmt.Println("Suma: ", result)

	//no es necesario poner los dos puntos porque ya se definio en la suma
	result = a - b

	fmt.Println(result)

	result = a * b

	fmt.Println(result)

	result = a / b

	fmt.Println(result)
}
