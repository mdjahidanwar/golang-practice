package main 

import (
"fmt"

)

func main() {
 

 fmt.Print("Please print your age: ")
 var age int
 fmt.Scanf("%d",&age)
/*
 if age < 3 {
	fmt.Println("infant")
}else if age >2 && age <13 {
	fmt.Println ("children") 


} else if age > 12 && age <= 19 {
	fmt.Println("teenage")


}else {
	fmt.Println("adult")

}
*/

switch age {
	case 2:
            fmt.Println("infant")
	case 3,4,5,6,7,8,9,10,11,12:
            fmt.Println("children")
	case 13,14,15,16,17,18,19:
            fmt.Println("teanage")
	default:
		fmt.Println("adult")
}



 //fmt.Println(age)


}