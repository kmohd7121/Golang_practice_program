// check number is palindrom

package main

import "fmt"
func main(){
	fmt.Println("enter a number\n")
	var x int 
	fmt.Scanf("%d",&x)
	check(x)
	
	
}
func check(x int)int {
	var m int
	m=x
	r:=0
	rev:=0
	for x!=0{
		r=x%10
		rev=rev*10+r
		x=x/10
	}
	if rev==m{
		fmt.Println("number is palindrom number ",m)
	}else{
		fmt.Println("number is not palindrom number ",m)
	}
	return rev
}