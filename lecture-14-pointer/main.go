package main
import "fmt"

func update (a *int ) {
fmt.Println(a)
*a=*a+10
}


func main() {
var x int 
var y *int
fmt.Println("value of x is: ",x)
fmt.Println("memory address of x is: ",&x)
fmt.Println ("value of y is: ",y)
fmt.Println("memory address of y is: ",&y)
x=10 //assignment 
y=&x //referencing
fmt.Println("value of x is: ",x) //accessing 
fmt.Println("value of y is: ",y)
fmt.Println("value of y is: ",*y) //dereferencing

update (&x)
fmt.Println(x)


}