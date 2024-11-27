package pkg_test

import (
	"math"
	"slices"
	"testing"

	"github.com/davidkleiven/hmatch/pkg"
)

func TestPeakDetection(t *testing.T) {
	data := make([]float64, 128)

	for i := range data {
		data[i] = math.Sin(12.0 * math.Pi * float64(i) / 128.0)
	}

	peaks := pkg.DetectSpectrumPeaks(data, 1.0, 1024)

	significantPeaks := slices.DeleteFunc(peaks, func(peak pkg.Peak) bool { return peak.Value < 1.0 })
	if len(significantPeaks) != 1 {
		t.Errorf("Perfect sinusoidal signal should only have one peak")
	}

	if significantPeaks[0].Frequency != 6.0/128.0 {
		t.Errorf("Expected position to be 1 got %f", significantPeaks[0].Frequency)
	}
}
