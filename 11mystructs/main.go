package main

import "fmt"

func main()  {
	fmt.Println("Structs in golang")
	// no inheritence in golang; No super or parent

	pratik := User{"Pratik", "pratik1971sinha@gmail.com", true, 21}
	fmt.Println(pratik)
	fmt.Printf("Pratik details are: %+v\n", pratik)
	fmt.Printf("Name is %v and Email is %v.: ", pratik.Name, pratik.Email)

}

type User struct {
	Name	string
	Email	string
	Status	bool
	Age		int 
}


