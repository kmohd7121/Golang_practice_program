// reverse number using function
package main

import "fmt"
func main(){
	fmt.Println("enter a number\n")
	var x int 
	fmt.Scanf("%d",&x)
	z:=check(x)
	fmt.Println("reverse of number  is",z)
	
}
func check(x int)int {
	r:=0
	rev:=0
	for x!=0{
		r=x%10
		rev=rev*10+r
		x=x/10
	}
	return rev
}