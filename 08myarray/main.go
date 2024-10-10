package main

import "fmt"

func main()  {
	fmt.Println("Welcome to array in Golang")
	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[2] = "Mango"
	fruitlist[3] = "Orange"

	fmt.Println("fruit list is: ", fruitlist)
	fmt.Println("fruit list length is: ", len(fruitlist))

	var veglist = [3]string{"Potato", "beans", "tomato"}
	fmt.Println("Veg list is: ", len(veglist))


}