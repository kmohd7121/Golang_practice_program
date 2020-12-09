//Program to print the element of an array present on odd position

package main
import "fmt"
func main(){
	arr:=[]int {1,2,3,4,5,6,7,8,9}
	fmt.Println("Original array")
	fmt.Println(arr)
	fmt.Println("Element of an array present on odd position")
	for i:=1;i<len(arr);i=i+2{
		fmt.Println(arr[i])
	}
}