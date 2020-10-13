package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens(t *testing.T) {
	cases := []struct {
		Input  int
		Expect [][]string
	}{
		{4, [][]string{}},
		{2, [][]string{}},
	}
	for _, caseItem := range cases {
		result := solveNQueens(caseItem.Input)
		assert.Equal(t, caseItem.Expect, result)
	}
}
