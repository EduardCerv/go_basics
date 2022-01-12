package main

import "fmt"

func main() {
	hello := "hello"
	world := "world"
	name := "Eduardo"
	age := 30
	//Println prints each new line with the message
	fmt.Println(hello)
	fmt.Println(world)

	fmt.Printf("Hello, %s and your age is %d \n", name, age)
	// %v is used when we don't know the kind of data to print
	fmt.Printf("Hello, %v and your age is %v \n", name, age)

	// Sprintf format the string and return the value to stored in a variable
	message := fmt.Sprintf("Hola %s, tu edad es: %d", name, age)
	fmt.Println(message)

	// use %T to know the kind of the variable is
	fmt.Printf("Kind of variable: %T \n", name) // prints 'Kind of variable: string'

	fmt.Print("Type a name:")
	// & is used for re-assign the value from the console
	fmt.Scanln(&name)

	fmt.Println(name)
}
