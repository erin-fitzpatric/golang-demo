package main

import "fmt"

func main() {
	// arrays - length cannot change
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"boo", "finn", "rey", "katie"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85) // must store the value in a var after appending to slices

	fmt.Println(scores, len(scores))

	// slice ranges - length can change
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "taco")
	fmt.Println(rangeOne)
}
