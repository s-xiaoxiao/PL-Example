package main

import "fmt"

func main() {
	fmt.Println("if else grammar")

	if 8%2 == 0 {
		fmt.Println("8 is even")
	} else {
		fmt.Println("8 is odd")
	}

	if 8%2 == 0 {
		// /də'vɪzəbl/
		fmt.Println("8 is divisible by 2")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative	")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
