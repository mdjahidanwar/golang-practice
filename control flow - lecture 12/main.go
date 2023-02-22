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
}else if age >2 && age <13 {
	fmt.Println ("children") 


} else if age > 12 && age <= 19 {
	fmt.Println("teenage")


}else {
	fmt.Println("adult")

}


 //fmt.Println(age)


}