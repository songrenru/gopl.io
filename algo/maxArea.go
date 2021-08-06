// m1:暴力破解
// m2:双指针
package main

import "fmt"

// func maxArea(height []int) int {
//     max := 0
//     i, j := 0, len(height)-1
//     for i < j {
//     	l, r := height[i], height[j]
// 		if cap := min(l, r) * (j - i); cap > max {
// 			max = cap
// 		}
// 		if l < r {
// 			i++
// 		} else {
// 			j--
// 		}
//     }
//     return max
// }

// func min(i, j int) int {
// 	if i < j {
// 		return i
// 	}
// 	return j
// }

// func maxArea(height []int) int {
// 	max := 0
// 	i, j := 0, len(height) - 1
// 	for i < j {
// 		if cap := ((j - i) * min(height[i], height[j])) > max {
// 			max = cap
// 		}
// 		if height[i] < height[j] {
// 			i++
// 		} else {
// 			j--
// 		}
// 	}
// 	return max
// }

func maxArea(height []int) int {
	i, j := 0, len(height) - 1
	max := (j - i) * min(height[i], height[j])
	for i < j {
		if height[i] < height[j] {
			if height[i] < height[i+1] {
				if cap := ((j - i) * min(height[i], height[j])); cap > max {
					max = cap
				}
			}
			i++
		} else {
			if height[j-1] < height[j] {
				if cap := ((j - i) * min(height[i], height[j])); cap > max {
					max = cap
				}
			}
			j--
		}
	}
	return max
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	res := maxArea(height)
	fmt.Println(res)
}