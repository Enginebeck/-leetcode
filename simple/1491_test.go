package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	cases := []struct {
		Input  []int
		Output float64
	}{{
		[]int{4000, 3000, 1000, 2000},
		2500.00000,
	}, {
		[]int{1000, 2000, 3000},
		2000.00000,
	}}
	for _, caseItem := range cases {
		result := average(caseItem.Input)
		assert.EqualValues(t, caseItem.Output, result)
	}
}
