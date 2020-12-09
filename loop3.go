// print table 
package main
import "fmt"
func main(){
	var x int 
	var table int 
	fmt.Println("enter a number \n")
	fmt.Scanf("%d\n",&x)
	for i:=1;i<=10;i++{
		table=i*x
		fmt.Println(table)
	}
	
}