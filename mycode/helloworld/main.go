package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
	name, marks, age := twoReturns("Sayan")
	fmt.Println(name, marks, age)
}

func twoReturns(name string) (string, int, int) {
	return name, 10, 5
}
