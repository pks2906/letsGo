package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to slices in Golang")

	var fruitlist = []string{"Apple", "Mango", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	 fruitlist = append(fruitlist, "Tomato", "Banana")
	 fmt.Println(fruitlist)

	 fruitlist = append(fruitlist[1:])
	 fmt.Println(fruitlist)

	 highScores := make([]int, 4)

	 highScores[0] = 234
	 highScores[1] = 945
	 highScores[2] = 465
	 highScores[3] = 867
	 //highScores[4] = 777

	 highScores = append(highScores, 555, 666, 321)

	 fmt.Println(highScores)

	 sort.Ints(highScores)
	 fmt.Println(highScores)
	 fmt.Println(sort.IntsAreSorted(highScores))

	 //how to remove a value from slices based on index

	 var courses = []string{"React", "JS", "Python", "C++", "Ruby"}
	 fmt.Println(courses)
	 var index int = 2
	 courses = append(courses[:index], courses[index+1:]...)
	 fmt.Println(courses)
}