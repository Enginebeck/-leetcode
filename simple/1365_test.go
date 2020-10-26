package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallerNumbersThanCurrent(t *testing.T) {
	cases := []struct {
		Input  []int
		Output []int
	}{{
		[]int{8, 1, 2, 2, 3},
		[]int{4, 0, 1, 1, 3},
	}, {
		[]int{6, 5, 4, 8},
		[]int{2, 1, 0, 3},
	}, {
		[]int{3, 3, 3, 3},
		[]int{0, 0, 0, 0},
	}}
	for _, caseItem := range cases {
		result := smallerNumbersThanCurrent(caseItem.Input)
		assert.Equal(t, caseItem.Output, result)
	}
}
