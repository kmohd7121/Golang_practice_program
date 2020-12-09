// check number is armstrong or not

// check number is palindrom

package main

import "fmt"
func main(){
	fmt.Println("enter a number\n")
	var x int 
	fmt.Scanf("%d",&x)
	check(x)
	
	
}
func check(x int) {
	var m int
	m=x
	r:=0
	arm:=0
	for x!=0{
		r=x%10
		arm=arm+r*r*r
		x=x/10
	}
	if arm==m{
		fmt.Println("number is armstrong number ",m)
	}else{
		fmt.Println("number is not armstrong number ",m)
	}

}