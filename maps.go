package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"tofee pudding": 2.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		23453425:  "mario",
		234324:    "luigi",
		234823904: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[23453425])

	phonebook[23453425] = "bowser"
	fmt.Println(phonebook[23453425])

	phonebook[234823904] = "yoshi"
	fmt.Println(phonebook)
}
