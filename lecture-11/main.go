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
fmt.Println(b1)
}