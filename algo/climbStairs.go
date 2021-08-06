// m1:recursive
// m2:dynamic programming
package main

import "fmt"

// func climbStairs(n int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

func climbStairs(n int) int {
	p, q, r := 0, 0 ,1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q	
	}
	return r
}

func main() {
	fmt.Println(climbStairs(10))
}