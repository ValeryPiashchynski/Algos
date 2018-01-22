package main

import (
	"fmt"

)

func main() {
	fmt.Println(singleNumber([]int{17, 12, 5, -6, 12, 4, 17, -5, 2, -3, 2, 4, 5, 16, -3, -4, 15, 15, -4, -5, -6}))
}

func singleNumber(nums []int) int {
	//My solution
	//sort.Ints(nums)
	//if len(nums) < 2 {
	//	return nums[0]
	//}
	//
	//for i := 0; i < len(nums)-1; i += 2 {
	//	if nums[i] != nums[i+1] {
	//		return nums[i]
	//	}
	//}
	//return nums[len(nums)-1]

	//Solution from leetcode
	a := 0
	for _, num := range nums {
		a ^= num
	}

	return a
}
