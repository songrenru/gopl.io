package main

import "fmt"

const num = 4

func main() {
	arr := [num]int{}
	num := 8 // 这种写法，会导致局部变量覆盖包级常量，导致runtime error
	for i := 0; i < num; i++ {
		fmt.Println(arr[i])
	}
}