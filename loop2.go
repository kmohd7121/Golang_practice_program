//sum of natural number

package main

import "fmt"
func main(){
	var x int 
	sum:=0
	fmt.Print("enter a number\n")
	fmt.Scanf("%d",&x)
	for i:=1;i<=x;i++{
		sum=sum+i
	}
	fmt.Printf("%d",sum)
}