package simple

import "fmt"

func rotate(nums []int, k int) []int {
	strictSteps := k % len(nums)
	copy(nums, append(nums[len(nums)-strictSteps:], nums[:len(nums)-strictSteps]...))
	fmt.Println(nums)
	return nums
}
