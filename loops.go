package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
		i++
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for in
	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	// use _ if you don't need to use index or value
	for _, value := range names {
		fmt.Printf("the value at index is %v \n", value)
	}

	for _, value := range names {
		fmt.Printf("the value at index is %v \n", value)
		value = "new string" // does not alter original slice
	}

	fmt.Println(names)
}
