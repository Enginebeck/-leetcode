package simple

import (
	"fmt"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	cases := []struct {
		Input  int
		Expect int
	}{
		{5, 1702766719},
	}
	for _, caseItem := range cases {
		result := firstBadVersion(caseItem.Input)
		fmt.Println(result, caseItem.Expect)
	}
}
