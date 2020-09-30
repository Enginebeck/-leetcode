package simple

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%3 != 0 {
		return false
	}
	if n == 3 {
		return true
	}
	if n > 3 {
		return isPowerOfThree(n / 3)
	}
	return false
}
