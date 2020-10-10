package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		Input  []int
		Target int
		Expect []int
	}{
		{
			Input:  []int{123},
			Target: 111,
			Expect: []int{123},
		},
	}
	for _, testCase := range cases {
		result := searchRange(testCase.Input, testCase.Target)
		assert.Equal(t, testCase.Expect, result)
	}
}
