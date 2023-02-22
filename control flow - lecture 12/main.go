package main 

import (
"fmt"

)

func main() {
 fmt.Print("Please print your age: ")
 var age int
 fmt.Scanf("%d",&age)

 if age < 3 {
	fmt.Println("infant")
}

 //fmt.Println(age)


}