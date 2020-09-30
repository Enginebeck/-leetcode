package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestUnivaluePath(t *testing.T) {
	cases := []struct {
		Input  *TreeNode
		Expect bool
	}{{}}
	for _, caseItem := range cases {
		result := longestUnivaluePath(caseItem.Input)
		assert.Equal(t, caseItem.Expect, result)
	}
}
