// copy one array element to another 
package main

import "fmt"

func main(){
	arr1:=[]int{1,2,3,4,5}
	arr2:=[]int {}
	arr2 = arr1
	fmt.Println("first array element\n")
	fmt.Println(arr1)
	fmt.Println("second array element\n")
	fmt.Println(arr2)
}