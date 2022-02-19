package main

import "fmt"

func main() {
	age := 35
	name := "erin"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println(("hello world!"))
	fmt.Println(("goodbye world!"))
	fmt.Println("my age is", age, "and my name is", name)

	// Formatted String - Printf && %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name) // outputs variable
	fmt.Printf("my age is %q and my name is %q \n", age, name) // adds quotes are string variables
	fmt.Printf("age is of type %T \n", age)                    // outputs type
	fmt.Printf("you scored %0.1f points! \n", 225.55)          // floats

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}
