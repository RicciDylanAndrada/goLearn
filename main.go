package main

import "fmt"

func main() {
	fmt.Println("hello world")

	fmt.Println("go" + "lang")
	fmt.Println("1+1", 1+1)
	fmt.Println((true && false))
	fmt.Println("7.0/3.0", 7.0/3.0)
	fmt.Println("true|false", true || false)
	fmt.Println(!true)

	fmt.Println("Array Examples")
	Arrays()

	fmt.Println("Functions Examples")
	Functions()
}
