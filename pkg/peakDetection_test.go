package pkg

import (
	"math"
	"slices"
	"testing"
)

func TestPeakDetection(t *testing.T) {
	data := make([]float64, 128)

	for i := range data {
		data[i] = math.Sin(12.0 * math.Pi * float64(i) / 128.0)
	}

	peaks := DetectSpectrumPeaks(data, 1.0, 1024)

	significantPeaks := slices.DeleteFunc(peaks, func(peak Peak) bool { return peak.Value < 1.0 })
	if len(significantPeaks) != 1 {
		t.Errorf("Perfect sinusoidal signal should only have one peak")
	}

	if significantPeaks[0].Frequency != 6.0/128.0 {
		t.Errorf("Expected position to be 1 got %f", significantPeaks[0].Frequency)
	}
}
