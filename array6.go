//Program to print the elements of an array present on even position
package main
import "fmt"
func main(){
	arr:=[]int {1,2,3,4,5,6,7,8,9}
	fmt.Println("Original array")
	fmt.Println(arr)
	fmt.Println("Element of an array present on even position")
	for i:=0;i<len(arr);i=i+2{
		fmt.Println(arr[i])
	}
}