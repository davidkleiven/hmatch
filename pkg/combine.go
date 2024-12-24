package pkg

import (
	"errors"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type DrawbarSetting [9]int

func Combine(positions []DrawbarSetting) DrawbarSetting {
	result := make([]int, 9)
	totalAmplitude := make([]float64, 9)
	for _, position := range positions {
		for i, v := range position {
			if v > 0 {
				totalAmplitude[i] += linearPower(v)
			}
		}
	}

	maxAllowedAmplitude := linearPower(8)
	maxAmplitude := slices.Max(totalAmplitude)

	if maxAmplitude > maxAllowedAmplitude {
		for i := range totalAmplitude {
			totalAmplitude[i] *= maxAllowedAmplitude / maxAmplitude
		}
	}

	for i, v := range totalAmplitude {
		result[i] = drawbarPosition(v)
	}
	return DrawbarSetting(result)
}

func linearPower(drawbarPosition int) float64 {
	intensity := 3.0 * (float64(drawbarPosition) - 1.0)
	return math.Pow(10.0, intensity/10.0)
}

func drawbarPosition(power float64) int {
	if power < 1e-6 {
		return 0
	}
	db := 10.0 * math.Log10(power)
	return 1 + int(math.Round(db/3.0))
}

func isOnlyDigitsAndCommas(input string) bool {
	re := regexp.MustCompile(`^[0-9,]+$`)
	return re.MatchString(input)
}

var ErrInvalidCharacters = errors.New("input contains other characters than digits and commas")
var ErrInvalidLengthOfSubtoken = errors.New("a subtoken does not contain 9 characters")

func ParseDrawbarPositions(positions string) ([]DrawbarSetting, error) {
	settings := []DrawbarSetting{}
	if !isOnlyDigitsAndCommas(positions) {
		return settings, ErrInvalidCharacters
	}

	for _, part := range strings.Split(positions, ",") {
		if len(part) != 9 {
			return settings, ErrInvalidLengthOfSubtoken
		}

		result := DrawbarSetting{}
		for i, char := range part {
			digit, err := strconv.Atoi(string(char))

			if err != nil {
				return settings, err
			}

			result[i] = digit
		}
		settings = append(settings, result)
	}

	return settings, nil
}
