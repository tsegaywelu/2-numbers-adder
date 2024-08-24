package main

import "fmt"

var show = fmt.Println
func Add(x,y int ) int {
	reult:=x+y
	return reult
}

func main() {

	show("plase enter two numbers to add ")
	var x, y int

	fmt.Scan(&x,&y)

	sum:=Add(x,y)
	show("the sum is :",sum)

}