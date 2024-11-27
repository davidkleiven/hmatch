package pkg

type StereoData struct {
	Left       []float64
	Right      []float64
	SampleRate float64
}

func NewStereoChannels(size int) StereoData {
	return StereoData{
		Left:  make([]float64, size),
		Right: make([]float64, size),
	}
}

func (s *StereoData) Mean() []float64 {
	result := make([]float64, len(s.Left))

	for i := range s.Left {
		result[i] = 0.5 * (s.Left[i] + s.Right[i])
	}
	return result
}
