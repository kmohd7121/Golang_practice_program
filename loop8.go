//sum of number 

package main

import "fmt"
func main(){
	fmt.Println("enter a number\n")
	var x int 
	fmt.Scanf("%d",&x)
	r:=0
	sum:=0
	for x!=0{
		r=x%10
		sum=sum+r
		x=x/10
	}
	fmt.Println("sum of number is",sum)
}