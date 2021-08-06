package main

import "fmt"

func main() {
	// var ages map[string]int // 只申明，则ages=nil，存入一个元素会报panic: assignment to entry in nil map
	var ages = map[string]int{}
	ages["key"] = 123
	fmt.Println("done")
}