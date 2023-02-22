package main
import "fmt"
//method 1
/*
func add(x int, y int) int {
r:=x+y
return r
}
*/

//method 2 
func add(x, y int) int {
r:=x+y
return r
}



func main (){

x:=add (10,30)
fmt.Println(x)
}