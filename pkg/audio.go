package pkg

import (
	"io"
	"os"

	"github.com/youpy/go-riff"
	"github.com/youpy/go-wav"
)

func Read(reader riff.RIFFReader) (StereoData, error) {
	wavReader := wav.NewReader(reader)
	format, err := wavReader.Format()

	if err != nil {
		return StereoData{}, err
	}

	result := StereoData{}
	result.SampleRate = float64(format.SampleRate)
	for {
		samples, err := wavReader.ReadSamples()

		if err == io.EOF {
			break
		}

		if err != nil {
			return StereoData{}, err
		}

		newLeft := make([]float64, len(samples))
		newRight := make([]float64, len(samples))

		for i, value := range samples {
			newLeft[i] = float64(wavReader.IntValue(value, 0))
			newRight[i] = float64(wavReader.IntValue(value, 1))
		}
		result.Left = append(result.Left, newLeft...)
		result.Right = append(result.Right, newRight...)
	}
	return result, nil
}

func ReadFromFile(fname string) (StereoData, error) {
	file, err := os.Open(fname)

	if err != nil {
		return StereoData{}, err
	}
	defer file.Close()
	return Read(file)
}
