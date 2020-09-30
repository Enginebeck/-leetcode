package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		Input  []int
		Target int
		Expect []int
	}{
		{
			Input:  []int{3, 2, 4},
			Target: 6,
			Expect: []int{1, 2},
		},
		{
			Input:  []int{2, 7, 11, 15},
			Target: 9,
			Expect: []int{0, 1},
		},
	}
	for _, testCase := range cases {
		result := twoSum(testCase.Input, testCase.Target)
		assert.Equal(t, testCase.Expect, result)
	}
}
