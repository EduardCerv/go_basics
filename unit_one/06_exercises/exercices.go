package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func quotient(a, b int) int {
	return a / b
}

func residue(a, b int) int {
	return a % b
}


func main() {
	a := 0
	b := 0

	fmt.Println("Type a number:")
	fmt.Scanln(&a)

	fmt.Println("Type another number:")
	fmt.Scanln(&b)

	res := sum(a, b)

	fmt.Println("The sum is:", res)

	quotient := quotient(a, b)
	residue := residue(a, b)

	fmt.Println("The quotient is:", quotient)
	fmt.Println("The residue is:", residue)
}