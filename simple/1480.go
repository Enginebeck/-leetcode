package simple

func runningSum(nums []int) []int {
	var result []int
	for index, item := range nums {
		prev := 0
		if index > 0 {
			prev = result[index-1]
		}
		result = append(result, item+prev)
	}
	return result
}
