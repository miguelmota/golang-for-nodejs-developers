package example

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	for _, tt := range []struct {
		a   int
		b   int
		ret int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{5, 5, 10},
	} {
		t.Run(fmt.Sprintf("(%v + %v)", tt.a, tt.b), func(t *testing.T) {
			ret := sum(tt.a, tt.b)
			if ret != tt.ret {
				t.Errorf("want %v, got %v", tt.ret, ret)
			}
		})
	}
}

func sum(a, b int) int {
	return a + b
}
