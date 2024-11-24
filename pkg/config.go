package pkg

func drawbarFeets() []float64 {
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
	feets := drawbarFeets()
	frequencies := make([]float64, len(feets))
	for i, feet := range feets {
		frequencies[i] = fundamentalFrequency * feets[fundamentalDrawbar] / feet
	}
	return frequencies
}
