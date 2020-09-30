package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfThree(t *testing.T) {
	cases := []struct {
		Input  int
		Expect bool
	}{
		{3, true},
		{4, false},
	}
	for _, caseItem := range cases {
		result := isPowerOfThree(caseItem.Input)
		assert.Equal(t, caseItem.Expect, result)
	}
}
