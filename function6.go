// check number is prime or not prime

package main
import "fmt"
func main(){
	fmt.Println("enter a number to check prime or not")
	var x int 
	fmt.Scan(&x)
	prime(x)
}
func prime(x int){
	var i int
	i=0
	if x%2!=0{
		i++
	}
	if(i==1){
		fmt.Println("number is prime ")
	}else{
		fmt.Println("number is not prime")
	}
}