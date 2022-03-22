package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, value := range a {
		fmt.Println(i, ":", value)
	}
	b := 0
	for b < 10 {
		fmt.Println(b)
		b++
	}
}
