package main

import "fmt"

func main() {
	var username string = "Pratik"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallfloat float64 = 255.4555734583685763587
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	//default values and aliases

	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type: %T \n", anothervariable)

	// implcit type 

	var website = "google.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000
	fmt.Println(numberOfUser)
}

