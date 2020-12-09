// print element of array
package main
import "fmt"
func main(){
	arr:=[3]int {1,2,3}
	for i:=0;i<len(arr);i++{
		fmt.Println("element of array ",arr[i])
	}
}