package simple

func twoSum(nums []int, target int) []int {
	var result []int
	for index, num := range nums {
		childNums := nums[index+1:]
		for childIndex, childNum := range childNums {
			if num+childNum == target {
				result = append(result, index, childIndex+index+1)
			}
		}
		if len(result) > 0 {

			break
		}
	}
	return result
}
