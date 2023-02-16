package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello bangladesh")
	// fmt.Printf("%c\n", 65)
	// fmt.Printf("%c\n", 0x41)
	// fmt.Printf("%d\n", 'অ')
	// fmt.Printf("%x\n", 'অ')
  numToChar(65)
  add(1,5)
}

func numToChar(num int ) {
 fmt.Printf("%c\n", num)
}

func add(num1 int, num2 int) {
	fmt.Println(num1+num2)
}