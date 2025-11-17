package euler

import (
	"testing"
)

func TestEuler(t *testing.T) {
	sum := 0
	max := 260000
	for i := 1; i <= max; i++ {
		square := i * i
		if square%2 != 0 {
			sum += square
		}
	}
	t.Logf("%v\n", sum)
}
