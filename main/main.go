package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	fmt.Println(a)
	a[4] = 2
	fmt.Println(a)
	z := a
	fmt.Println(z)
	z[2] = 3
	fmt.Println(z)
	fmt.Println(&a)
	a = make([]int, 6)
	fmt.Println(a)
}
