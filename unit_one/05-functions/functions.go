package main

import "fmt"

func greeting(name string) {
	//fmt.Println("Hello world")
	fmt.Println("Hello,", name)
}

//if you pass same kind of data just need to defined by the las parameter
//also we need to defined the kind of data returned
func sum(a, b int) int {
	return a + b
}

func main() {
	greeting("Eduardo")
	res := sum(127, 127)
	fmt.Println("The result is:", res)
}
