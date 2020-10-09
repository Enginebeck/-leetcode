package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	cases := []struct {
		Input  []int
		Output []float64
	}{{
		[]int{1, 2, 3, 4},
		[]float64{1, 3, 6, 10},
	}}
	for _, caseItem := range cases {
		result := average(caseItem.Input)
		assert.EqualValues(t, caseItem.Output, result)
	}
}
