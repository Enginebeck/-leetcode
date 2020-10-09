package simple

func reverseString(s []byte) {
	strLength := len(s)
	for index := 0; index < strLength/2; index++ {
		s[index], s[strLength-1-index] = s[strLength-1-index], s[index]
	}
}
