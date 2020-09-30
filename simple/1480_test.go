package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningSum(t *testing.T) {
	cases := []struct {
		Input  []int
		Output []int
	}{{
		[]int{1, 2, 3, 4},
		[]int{1, 3, 6, 10},
	}}
	for _, caseItem := range cases {
		result := runningSum(caseItem.Input)
		assert.Equal(t, caseItem.Output, result)
	}
}
