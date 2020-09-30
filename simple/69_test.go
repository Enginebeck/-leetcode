package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	cases := []struct {
		Input  int
		Expect int
	}{
		{4, 2},
		{1, 1},
	}
	for _, caseItem := range cases {
		result := mySqrt(caseItem.Input)
		assert.Equal(t, caseItem.Expect, result)
	}
}
