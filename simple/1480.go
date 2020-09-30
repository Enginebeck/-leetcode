package simple

func runningSum(nums []int) []int {
	for index, _ := range nums {
		if index > 0 {
			nums[index] += nums[index-1]
		}
	}
	return nums
}
