package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		Input  []string
		Expect string
	}{
		{[]string{"lee", "leet"}, "lee"},
		{[]string{"lt", "hi"}, ""},
	}
	for _, caseItem := range cases {
		result := longestCommonPrefix(caseItem.Input)
		assert.Equal(t, caseItem.Expect, result)
	}
}
