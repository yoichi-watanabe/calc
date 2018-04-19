package calc

import "testing"

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Fatal("sum(1, 2) shold be 3, but doesn't match")
	}
}
