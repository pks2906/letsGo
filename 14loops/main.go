package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d:=0; d< len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	//super imp

	// for index, day := range days{
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rouguevalue := 1
	//while of other language
	for rouguevalue < 10 {

		if rouguevalue == 2 {
			goto lco
		}

		if rouguevalue == 5 {
			// break
			rouguevalue++
			continue
		}
		fmt.Println("Value is: ", rouguevalue)
		rouguevalue++
	}
	//goto statement
lco:
	fmt.Println("Jumping on codeonline.in")
}
