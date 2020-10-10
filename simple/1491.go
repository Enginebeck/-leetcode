package simple

func average(salary []int) float64 {
	max, min := salary[0], salary[0]
	total := 0
	for _, val := range salary {
		if max < val {
			max = val
		} else if min > val {
			min = val
		}
		total += val
	}
	return float64(total-max-min) / (float64(len(salary) - 2))
}
