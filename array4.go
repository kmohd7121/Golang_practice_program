// Program to print duplicate element of an array
// Program to print duplicate element of an array

package main

import "fmt"

func main() {
	arr := []int {1, 2, 3, 4, 2, 7, 8, 8, 3}
	fmt.Println("Duplicate element in give array\n")
	for i:= 0;i<len(arr);i++ {
		for j:=i+1;j<len(arr);j++ {
			if(arr[i] == arr[j]) {
				fmt.Printf(" %d \n", arr[j])
			}
		}
	}
}
