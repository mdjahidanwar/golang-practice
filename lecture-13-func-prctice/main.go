package main 
import "fmt"

//exmple- 1
/*
func add(x int, y int) int {
r:=x+y
return r
}
*/

//exmple-2
/*
func add(x, y int) int {
r:=x+y
return r
}
*/
//exmple-3 
/*
func add(x, y int) (r int) {
r=x+y
return r
}
*/
/*
//exmple-4
func add(x, y int) (r int) {
r=x+y
return
}

func rectangle(l int, b int) (area int, parameter int) {
parameter = 2* (l+b)
area = l*b
return 
}
*/

//pointer 

func update (a *int, t *string) {
*a = *a +5
*t = *t+" Doe"
return
}

func main(){
//x:=add(10,30)
//fmt.Println(x)
//p,k:=rectangle(12,12)
//fmt.Println(p,k)

a:=10
t:= "mostain"
update(&a,&t)
fmt.Println(a,t)
}