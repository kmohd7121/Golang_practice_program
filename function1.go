// add two number using function
package main
import "fmt"
func main(){
	var x int 
	var y int 
	fmt.Println("enter two number\n")
	fmt.Scanf("%d %d",&x,&y)
	z:=add(x,y)
	fmt.Printf("%v",z)
}
func add(x int , y int)int {
	sum:=0
	sum=x+y
	return sum
}