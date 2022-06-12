package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	for i, el := range nums {
		fmt.Println(i)
		fmt.Println(el)
	}
}
