package main
import "fmt"
//example 1
/*
func add(x int, y int) int {
r:=x+y
return r
}
*/

//example 2 
/*
func add(x, y int) int {
r:=x+y
return r
}
*/


//example 3

func add(x int, y int) (r int) {
r=x+y
return r
}



func main (){

x:=add (10,30)
fmt.Println(x)
}