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
/*
func add(x int, y int) (r int) {
r=x+y
return r
}
*/

//example 4
/*
func add(x int, y int) (r int) {
	r = x + y
	return
}
*/
//example 5 
func rectangle (l,b int) (area, parameter int ){
parameter = 2*(l+b)
area=l*b
return 
}
//example 6

func Add(x int, y int) (r int) {
r=x*y
return
}

func main() {

//	x := add(10, 30)
//	fmt.Println(x)
a,p:= rectangle (10,10)
fmt.Println(a,p)
c:=Add(10,10)
fmt.Println(c)
}
