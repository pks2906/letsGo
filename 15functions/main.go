package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions is golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := ProAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro Message is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func ProAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hello pro result function"
}

func greeter() {
	fmt.Println("Hello from golang")
}
