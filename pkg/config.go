package pkg

import (
	"errors"
	"strconv"

	"github.com/spf13/pflag"
)

func DrawbarFeets() []float64 {
	return []float64{
		16.0, 5.0 + 1.0/3.0, 8.0, 4.0, 2.0 + 2.0/3.0, 2.0, 1.0 + 3.0/5.0, 1.0 + 1.0/3.0, 1.0,
	}
}

type Fundamental int

const (
	Sixteen Fundamental = 0
	Eight               = 2
	Four                = 3
)

func drawbarFrequencies(fundamentalFrequency float64, fundamentalDrawbar int) []float64 {
	feets := DrawbarFeets()
	frequencies := make([]float64, len(feets))
	for i, feet := range feets {
		frequencies[i] = fundamentalFrequency * feets[fundamentalDrawbar] / feet
	}
	return frequencies
}

func IntFromFlag(flag *pflag.Flag) (int, error) {
	if flag == nil {
		return 0, errors.New("unknown flag")
	}

	return strconv.Atoi(flag.Value.String())
}
