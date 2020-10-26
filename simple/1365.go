package simple

func smallerNumbersThanCurrent(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, numsLen)
	for index, item := range nums {
		for subIndex := index + 1; subIndex < numsLen; subIndex++ {
			if item > nums[subIndex] {
				result[index]++
			}
			if nums[subIndex] > item {
				result[subIndex]++
			}
		}
	}
	return result
}
