// Program to short the element of an array in ascending order

package main
import "fmt"
func main(){
	arr:=[]int {12,34,54,2,35,6,75,98,4,5,}
	fmt.Println("Original array")
	for i:=0;i<len(arr);i++{
		fmt.Printf("%d\n",arr[i])
	}
	fmt.Println("shorted array ")
	temp:=0
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			if(arr[i]>arr[j]){
				temp=arr[i]
				arr[i]=arr[j]
				arr[j]=temp
			}
		}
	}
	fmt.Println("\n")
	for k:=0;k<len(arr);k++{
		fmt.Println(arr[k])
	}
}