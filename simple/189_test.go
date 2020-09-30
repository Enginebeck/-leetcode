package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		Input  []int
		Step   int
		Expect []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	}
	for _, caseItem := range cases {
		result := rotate(caseItem.Input, caseItem.Step)
		assert.Equal(t, caseItem.Expect, result)
	}
}
