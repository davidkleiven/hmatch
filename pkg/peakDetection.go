package pkg

import (
	"cmp"
	"math/cmplx"
	"slices"

	"github.com/davidkleiven/gosfft/sfft"
)

type Peak struct {
	Value    float64
	Position int
}

func DetectSpectrumPeaks(data []float64, relativeThreshold float64) []Peak {
	fft := sfft.NewFFT1(len(data))
	ftData := fft.FFT(data)
	amplitudes := make([]float64, len(ftData))

	for i, v := range ftData {
		amplitudes[i] = cmplx.Abs(v)
	}

	peaks := []Peak{}

	for i := 1; i < len(amplitudes)-1; i++ {
		if (amplitudes[i-1] < amplitudes[i]) && (amplitudes[i+1] < amplitudes[i]) {
			peaks = append(peaks, Peak{Value: amplitudes[i], Position: i})
		}
	}

	// Filter very small peaks
	maxValue := slices.MaxFunc(peaks, func(p1 Peak, p2 Peak) int { return cmp.Compare(p1.Value, p2.Value) })
	return slices.DeleteFunc(peaks, func(p Peak) bool { return p.Value < relativeThreshold*maxValue.Value })
}
