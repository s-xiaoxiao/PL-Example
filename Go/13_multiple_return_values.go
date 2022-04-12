package main

import "fmt"

func vals() (int, int) {
	return 3, 4
}
func main() {

	// Go 原生支持 多返回值。

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
