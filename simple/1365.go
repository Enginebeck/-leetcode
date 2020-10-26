package simple

func smallerNumbersThanCurrent(nums []int) []int {
	numsLen := len(nums)
	var result []int
	for index := 0; index < numsLen; index++ {
		result = append(result, 0)
	}
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
