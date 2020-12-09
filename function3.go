// area of rectangle using function

package main

import "fmt"

func main(){
	var len int 
	var width int
	fmt.Println("enter length ")
	fmt.Scan(&len)
	fmt.Println("enter width ")
	fmt.Scan(&width)
	 y:=area(len,width)
	fmt.Println("area of rectangle is",y)
}
func area(x int ,y int) int{
	z:=x*y
	return z
}