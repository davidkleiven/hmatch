package pkg_test

import (
	"math"
	"testing"

	"github.com/davidkleiven/hmatch/pkg"
)

func TestPeakDetection(t *testing.T) {
	data := make([]float64, 128)

	for i := range data {
		data[i] = math.Sin(2.0 * math.Pi * float64(i) / 128.0)
	}

	peaks := pkg.DetectSpectrumPeaks(data, 0.01)
	if len(peaks) != 1 {
		t.Errorf("Perfect sinusoidal signal should only have one peak")
	}

	if peaks[0].Position != 1 {
		t.Errorf("Expected position to be 1 got %d", peaks[0].Position)
	}
}
