package pkg_test

import (
	"slices"
	"testing"

	"github.com/davidkleiven/hmatch/pkg"
)

func TestGroupBy(t *testing.T) {
	intergers := []int{1, 2, 3, 4, 5}
	evenAndOddNumbers := pkg.Groupby(intergers, func(n int) int {
		return n % 2
	})

	expect := map[int][]int{
		0: {2, 4},
		1: {1, 3, 5},
	}

	if len(evenAndOddNumbers) != len(expect) {
		t.Errorf("Expected %v got %v", evenAndOddNumbers, expect)
	}

	for k, v := range expect {
		if val, ok := evenAndOddNumbers[k]; !ok || slices.Compare(val, v) != 0 {
			t.Errorf("Expected %v got %v", evenAndOddNumbers, expect)
		}
	}
}
