package simple

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	prefix := ""

	if len(strs) == 0 {
		return prefix
	}
	minStr := strs[0]
	for _, str := range strs {
		if len(str) < len(minStr) {
			minStr = str
		}
	}
	for index := 0; index < len(minStr); index++ {
		testPrefix := minStr[:len(minStr)-index]
		isPrefix := true
		for _, str := range strs {
			isPrefix = strings.HasPrefix(str, testPrefix)
			if !isPrefix {
				break
			}
		}
		if isPrefix {
			prefix = testPrefix
			break
		}
	}
	return prefix
}
