package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		Input  int
		Expect int
	}{
		{
			Input:  123,
			Expect: 321,
		},
		{
			Input:  -1,
			Expect: -1,
		},
		{
			Input:  1534236469,
			Expect: 0,
		},
	}
	for _, testCase := range cases {
		result := reverse(testCase.Input)
		assert.Equal(t, testCase.Expect, result)
	}
}
