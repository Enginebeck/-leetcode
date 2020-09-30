package simple

func mySqrt(x int) int {
	input := float64(x)
	result := 1.00
	precision := 1e-5
	for {
		temp := input - result*result
		if precision > temp && temp > -precision {
			break
		} else {
			result = (result + input/result) / 2
		}
	}
	if result < 0 {
		result = -result
	}
	return int(result)
}
