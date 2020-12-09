// swap two number without third variable

package main

import "fmt"

func main(){
	var x int
	var y int 
	fmt.Println("enter first number ")
	fmt.Scan(&x)
	fmt.Println("enter second number ")
	fmt.Scan(&y)
	fmt.Println("after swaping number")
	fmt.Println(swaping(x,y))
}
func swaping(x int ,y int)(int ,int){
	x=x+y
	y=x-y
	x=x-y
	return x,y

}