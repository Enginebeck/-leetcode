package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		Input  []byte
		Expect []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'g', 'o', 'o', 'd'}, []byte{'d', 'o', 'o', 'g'}},
		{[]byte{'h', 'w', 'o', 'd'}, []byte{'d', 'o', 'w', 'h'}},
		{
			[]byte{'A', ' ', 'm', 'a', 'n', ',', ' ', 'a', ' ', 'p', 'l', 'a', 'n', ',', ' ', 'a', ' ', 'c', 'a', 'n', 'a', 'l', ':', ' ', 'P', 'a', 'n', 'a', 'm', 'a'},
			[]byte{'a', 'm', 'a', 'n', 'a', 'P', ' ', ':', 'l', 'a', 'n', 'a', 'c', ' ', 'a', ' ', ',', 'n', 'a', 'l', 'p', ' ', 'a', ' ', ',', 'n', 'a', 'm', ' ', 'A'},
		},
	}
	for _, caseItem := range cases {
		reverseString(caseItem.Input)
		assert.Equal(t, caseItem.Expect, caseItem.Input)
	}
}
