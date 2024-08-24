package main

import (
	"fmt"

	adder "github.com/tsegaywelu/2-numbers-adder/adder"
)



func main() {

	fmt.Println("plase enter two numbers to add ")
	var x, y int

	fmt.Scan(&x,&y)

	sum:=adder.Adder(x,y)
	fmt.Println("the sum is :",sum)

}