package main

import (
	"fmt"
	"math/big"

	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to maths in golang")

// 	// var mynumberOne int = 2
// 	// var mynumbersTwo float64 = 4

// 	//fmt.Println("The sum is:", mynumberOne+int(mynumbersTwo))

	//random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}

