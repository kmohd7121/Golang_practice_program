//factorial of a number 

package main

import "fmt"

func main(){
	var num,factorial_val int 
	factorial_val=1
	fmt.Println("enter a number\n")
	fmt.Scanf("%d",&num)
	if num<0{
		fmt.Print("error negative number\n")

	}else if num==0{
		fmt.Println(1)
	}else{
		for i:=1;i<=num;i++{
			factorial_val=i*factorial_val
		}
		fmt.Printf("factorial of number is:%d",factorial_val)
	}
}