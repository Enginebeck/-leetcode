package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumGap(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect int
	}{
		{[]int{3, 6, 9, 1}, 3},
		{[]int{10}, 0},
	}
	for _, caseItem := range cases {
		result := maximumGap(caseItem.Input)
		assert.Equal(t, caseItem.Expect, result)
	}
}
