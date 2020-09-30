package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumGap(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect int
	}{
		{[]int{4}, 0},
		{[]int{1, 8}, 7},
		{[]int{1, 4, 8}, 4},
	}
	for _, caseItem := range cases {
		result := maximumGap(caseItem.Input)
		assert.Equal(t, caseItem.Expect, result)
	}
}
