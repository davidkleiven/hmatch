package pkg

import (
	"cmp"
	"math"
	"math/cmplx"
	"slices"

	"github.com/davidkleiven/gosfft/sfft"
	"gonum.org/v1/gonum/dsp/window"
)

type Peak struct {
	Identifier int
	Value      float64
	Frequency  float64
}

func DetectSpectrumPeaks(data []float64, sampleRate float64, hannWindowLength int) []Peak {
	N := slices.Min([]int{hannWindowLength, len(data)})
	amplitudes := make([]float64, N)

	fft := sfft.NewFFT1(N)
	for i := 0; i < int(len(data)/N); i++ {
		ftData := fft.FFT(window.Hann(data[i*N : (i+1)*N]))
		for j, v := range ftData {
			amplitudes[j] += math.Pow(cmplx.Abs(v), 2) / float64(N)
		}
	}

	peaks := []Peak{}
	peakCounter := 0
	for i := 1; i < len(amplitudes)-1; i++ {
		if (amplitudes[i-1] < amplitudes[i]) && (amplitudes[i+1] < amplitudes[i]) {
			freq := sampleRate * fft.Freq(i)
			peaks = append(peaks, Peak{Value: amplitudes[i], Frequency: freq, Identifier: peakCounter})
			peakCounter += 1
		}
	}
	return peaks
}

func maxPeak(peaks []Peak) Peak {
	return slices.MaxFunc(peaks, func(p1 Peak, p2 Peak) int { return cmp.Compare(p1.Value, p2.Value) })
}
