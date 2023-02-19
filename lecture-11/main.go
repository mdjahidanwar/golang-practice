package main
import (
"fmt"
)
//type book struct {
//title string 
//author string 
//pages int 
//ISBN string 
//price float32
//}

func main() {
//var b1 book

b1:= struct{
title string 
author string 
pages int
ISBN string 
price float32
}{
title: "an introduction",
author: "jahid",
ISBN: "9x9x9x9xx9",
price: 10.80,
pages: 10,
}


fmt.Println(b1)


var w1 weight 
w1=30.5
fmt.Println(w1)


}