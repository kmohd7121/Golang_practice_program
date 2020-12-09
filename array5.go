// Program to print the element of an array in reverse order

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original array")
	for i := 0; i < len(arr); i++ {
		fmt.Printf(" %d \n", arr[i])
	}
	fmt.Println("Reverse array")
	for j := len(arr)-1; j >=0; j-- {
		fmt.Printf("%d\n",arr[j])
	}
}
