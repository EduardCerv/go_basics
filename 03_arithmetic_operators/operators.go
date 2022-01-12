package main

import "fmt"

func main() {
	a := 20
	b := 10

	result := a + b

	fmt.Println("Sum: ", result)

	//we don't need to set the colon because it was assigned in the sum
	result = a - b

	fmt.Println(result)

	result = a * b

	fmt.Println(result)

	result = a / b

	fmt.Println(result)
}
