// find the greatest number in tha array
package main

import "fmt"

func main(){
	arr:=[]int {23,4,5,12,13,25,40}
	max:=arr[0]
	for i:=0;i<len(arr);i++{
		if arr[i]>max{
			max=arr[i]
		}
	}
	fmt.Println("greatest number in array",max)
}