package pkg

import (
	"math"
	"testing"
)

func TestDrawbarFrequency(t *testing.T) {
	frequencies := drawbarFrequencies(50.0, 0)

	if math.Abs(frequencies[2]/frequencies[0]-2.0) > 1e-6 {
		t.Errorf("Expected 8 feet to be twice the freqenccy of 16 feet. Got %v", frequencies)
	}
}
