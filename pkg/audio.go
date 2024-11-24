package pkg

import (
	"io"

	"github.com/youpy/go-riff"
	"github.com/youpy/go-wav"
)

func Read(reader riff.RIFFReader) ([]StereoData, error) {
	wavReader := wav.NewReader(reader)

	result := []StereoData{}
	for {
		samples, err := wavReader.ReadSamples()

		if err == io.EOF {
			break
		}

		if err != nil {
			return result, err
		}

		current := NewStereoChannels(len(samples))

		for i, value := range samples {
			current.Left[i] = float64(wavReader.IntValue(value, 0))
			current.Right[i] = float64(wavReader.IntValue(value, 1))
		}
		result = append(result, current)
	}
	return result, nil
}
