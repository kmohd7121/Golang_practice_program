// check arr is same or not
package main
import "fmt"
func main(){
	arr1:=[5]int{1,2,3,4,5}
	arr2:=[5]int{1,2,3,4,5}
	k:=0
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if arr1[i]==arr2[j]{
				k++;
			}else{
				continue
			}
		}
	}
	if k==5{
		fmt.Println("array element is same")
	}else{
		fmt.Println("array element is not same")
	}
}
