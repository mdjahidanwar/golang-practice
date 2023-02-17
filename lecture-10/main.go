package main
import (
"fmt"
//"reflect"
)

func main() {
a:=make([]string,0)
fmt.Println(a)

var fruits []string
fruits = append(fruits, "apple","banana")
fmt.Println(fruits)

fmt.Printf("%T",fruits)

}