package hard

func maximumGap(nums []int) int {
	maxGap := 0
	numsLen := len(nums)
	if numsLen < 2 {
		return maxGap
	}

	for index, _ := range nums {
		for internalCircleIndex := index + 1; internalCircleIndex < numsLen; internalCircleIndex++ {
			if nums[index] > nums[internalCircleIndex] {
				nums[index], nums[internalCircleIndex] = nums[internalCircleIndex], nums[index]
			}
		}
		if index > 0 {
			diff := nums[index] - nums[index-1]
			if diff > maxGap {
				maxGap = diff
			}
		}
	}
	return maxGap
}
