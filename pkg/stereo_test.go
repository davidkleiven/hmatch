package pkg

import (
	"testing"
)

func TestMean(t *testing.T) {
	data := StereoData{
		Left:  []float64{1.0, 2.0, 3.0},
		Right: []float64{2.0, 3.0, 4.0},
	}

	expect := []float64{1.5, 2.5, 3.5}

	result := data.Mean()

	for i := range expect {
		if (expect[i] - result[i]) > 1e-6 {
			t.Errorf("Expected %v got %v", expect, result)
		}
	}
}
