// math/math_test.go

package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum([]int{10, -2, 3})
	if sum != 11 {
		t.Errorf("fail want 11 bug got %d", sum)
	}
}
