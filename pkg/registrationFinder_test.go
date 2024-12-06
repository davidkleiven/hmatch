package pkg

import (
	"slices"
	"strings"
	"testing"
)

func TestFormatDrawbarPositions(t *testing.T) {
	positions := []int{0, 0, 8, 1, 2, 3, 2, 4, 5}

	expect := "00 8123 245"
	if result := FormatDrawbarPositions(positions); result != "00 8123 245" {
		t.Errorf("Got %s wanted %s", result, expect)
	}
}

func TestPeakToString(t *testing.T) {
	peaks := []Peak{
		{
			Identifier: 0,
			Value:      10.0,
			Frequency:  200.0,
		},
	}

	result := PeakToString(peaks)

	if !strings.Contains(result, "200.00 Hz") {
		t.Errorf("Expected 200.00 Hz to be part of %s", result)
	}
}

func TestRegistarationFinder(t *testing.T) {
	peaks := []Peak{
		{
			Identifier: 0,
			Value:      1e6,
			Frequency:  10.0,
		},
		{
			Identifier: 1,
			Value:      1e5,
			Frequency:  20.0,
		},
	}
	result := FindRegistration(peaks, 2)
	expectPositions := []int{0, 0, 8, 5, 0, 0, 0, 0, 0}

	if !slices.Equal(result.Positions, expectPositions) {
		t.Errorf("Got %v wanted %v", result.Positions, expectPositions)
	}
}
