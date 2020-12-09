//fabonacci number
package main
import "fmt"
func main(){
	var x int 
	fmt.Println("enter a number\n")
	fmt.Scanf("%d",&x)
	y:=0
	z:=1
	fmt.Printf("%d\n%d\n",y,z)
	for i:=1;i<=x-2;i++{
		n:=y+z
		fmt.Println(n)
		y=z
		z=n	
	}
} 