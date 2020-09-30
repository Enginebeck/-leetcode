package simple

import "fmt"

func runningSum(nums []int) []int {
	for index, _ := range nums {
		if index > 0 {
			nums[index] += nums[index-1]
		}
	}
	fmt.Println(nums)
	return nums
}
