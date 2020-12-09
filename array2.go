package main

import "fmt"

func main() {
	arr := []int{1, 2, 8, 3, 2, 2, 2, 5, 1}
	fr := [9]int{}
	visited := -1
	for i := 0; i < len(arr); i++ {
		count := 1
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count = count + 1
				fr[j] = visited
			}
		}
		if fr[i] != visited {
			fr[i] = count
		}
	}
	fmt.Println("------------------------\n")
	fmt.Println("Element | Frequency\n")
	fmt.Println("------------------------\n")
	for k := 0; k < len(arr); k++ {
		if fr[k] != visited {
			fmt.Printf(" %d || %d\n", arr[k], fr[k])
		}
	}
	fmt.Printf("--------------------\n")
}
