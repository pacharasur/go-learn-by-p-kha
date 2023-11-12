package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, world")
	// arr := []int{1, 2, 3, 4}
	arr := make([]int, 4)
	arr[0] = 30
	fmt.Println(arr)

	txt := "today is sunday"
	fmt.Println(txt[0:5])
	fmt.Println(arr[0:1])
	fmt.Println(len(txt))
}
