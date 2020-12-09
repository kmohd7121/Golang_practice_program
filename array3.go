// rotation of array

package main

import "fmt"

func main(){
	arr:=[]int{1,2,3,4,5}
	n:=3
	fmt.Println("-----original array-----")
	for i:=0;i<len(arr);i++{
		fmt.Printf(" %d ",arr[i])
	}
	//rotation the given array by n times oward left

	for i:=0;i<n;i++{
		var j int 
		var first int
		first=arr[0]
		for j=0;j<len(arr)-1;j++{
			arr[j]=arr[j+1]
		}
		arr[j]=first
	}
	fmt.Println("\n")
	fmt.Println("Array after left rotation\n")
	for i:=0;i<len(arr);i++{
		fmt.Printf(" %d ",arr[i])
	}
}