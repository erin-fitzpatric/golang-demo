package main

import "fmt"

// func updateName(n string) {
// 	n = "wedge"
// }

func updateName(n *string) {
	*n = "wedge"
}

func main() {
	name := "tifa"

	// updateName(name)

	// & gets the memory address of the value (pointer)
	// fmt.Println("memory address of name is:", &name)

	m := &name
	// fmt.Println("memory address:", m)
	// fmt.Println("value at memory address:", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}

/*
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|
*/
