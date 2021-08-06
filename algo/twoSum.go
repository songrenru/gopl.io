// m1:暴力破解
// m2:两遍循环
// m3:一遍循环
package main

import "fmt"

// func twoSum(nums []int, target int) []int {
//     dict := make(map[int]int)
//     for i, v := range nums {
//         nums[v] = i
//     }
//     for i, v := range nums {
//         val, ok := dict[target - nums[i]]
//         if ok && i != val {
//             return []int{i, val}
//         }
//     }
//     return nil
// }

func twoSum(nums []int, target int) []int {
    dict := make(map[int]int)
    for i, v := range nums {
        val, ok := dick[target - nums[i]]
        if ok {
            return []int{i, val}
        }
        dict[v] = i
    }
}

func main() {
	nums := []int{2,2,7,11,15}
	res :=twoSum(nums, 9)
	fmt.Println(res)
}