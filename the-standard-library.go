package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	// string operations do not alter the original string
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "there"))
	fmt.Println(strings.Split(greeting, "e"))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// sort will alter original slice
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) // SearchInts will return one more than length if not found
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
}
