package main
import "fmt"
func main(){
	fmt.Println("enter a number ")
	var x int
	fmt.Scan(&x)
	count(x)
}
func count(x int){
	i:=0
	for x!=0{
		x=x/10
		i=i+1
	}
	fmt.Println("total number of digit",i)
}s