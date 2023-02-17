package main
import (
"fmt"
)


type book struct{
title string 
author string 
ISBN string 
price float32
pages int
}

func main() {
var b1 book 
b1.title="an introduction to programming in go"
b1.author="caleb doxy"
b1.ISBN="978-1478355823"
b1.price=10.50
b1.pages=165
fmt.Println(b1)
}