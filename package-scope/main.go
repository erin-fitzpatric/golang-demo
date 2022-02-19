package main

import "fmt"

var score = 99.5

func main() {
	sayHello("mario")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}

// You have to "go run main.go greetings.go" to use both files
