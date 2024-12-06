package pkg

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"slices"
)

type DrawbarPositionResult struct {
	Positions  []int
	MaxFreqDev float64
}

type PeakAssignment struct {
	peak                Peak
	frequencyDifference float64
	drawbarNumber       int
}

func onlyKeepClosestDrawbarAssignmentPerPeak(assignments []PeakAssignment) []PeakAssignment {
	assignmentPerPeak := groupby(assignments, func(a PeakAssignment) int { return a.peak.Identifier })
	result := []PeakAssignment{}

	for _, a := range assignmentPerPeak {
		result = append(result, slices.MinFunc(a, func(a PeakAssignment, b PeakAssignment) int {
			return cmp.Compare(a.frequencyDifference, b.frequencyDifference)
		}))
	}
	return result
}

func PeakToString(peaks []Peak) string {
	result := "Detected peaks (<3000 Hz):\n"
	sortedPeaks := slices.SortedFunc(
		slices.Values(peaks),
		func(a Peak, b Peak) int { return cmp.Compare(a.Frequency, b.Frequency) },
	)
	for _, peak := range sortedPeaks {

		if peak.Frequency < 3000.0 {
			result += fmt.Sprintf("%8.2f Hz (amplitude: %1.2e)\n", peak.Frequency, peak.Value)
		}
	}
	return result
}

func FindRegistration(peaks []Peak, maxPeakDrawbarNo int) DrawbarPositionResult {
	maxPeak := maxPeak(peaks)
	frequencies := drawbarFrequencies(maxPeak.Frequency, maxPeakDrawbarNo)
	log.Printf("Peak frequence %.2f Hz", maxPeak.Frequency)

	closestPeaks := make([]PeakAssignment, len(frequencies))
	for i, freq := range frequencies {
		closest := slices.MinFunc(peaks, func(a Peak, b Peak) int {
			diffA := math.Abs(a.Frequency - freq)
			diffB := math.Abs(b.Frequency - freq)
			return cmp.Compare(diffA, diffB)
		},
		)
		closestPeaks[i] = PeakAssignment{
			peak:                closest,
			frequencyDifference: math.Abs(closest.Frequency - freq),
			drawbarNumber:       i,
		}
	}

	peakAssignments := onlyKeepClosestDrawbarAssignmentPerPeak(closestPeaks)
	positions := calculateDrawbarPositions(peakAssignments)

	freqDeviations := make([]float64, len(peakAssignments))
	for i, peak := range peakAssignments {
		freqDeviations[i] = peak.frequencyDifference
	}

	return DrawbarPositionResult{
		Positions:  positions,
		MaxFreqDev: slices.Max(freqDeviations),
	}
}

func calculateDrawbarPositions(assignPeaks []PeakAssignment) []int {
	desibelValues := make([]float64, len(assignPeaks))
	for i, v := range assignPeaks {
		desibelValues[i] = 10.0 * math.Log10(v.peak.Value)
	}

	numDrawbarPositions := len(DrawbarFeets())

	maxDb := slices.Max(desibelValues)

	drawbarPositions := make([]int, numDrawbarPositions)

	for i, db := range desibelValues {
		drawbarPositions[assignPeaks[i].drawbarNumber] = max(8-int(math.Round((maxDb-db)/3.0)), 0)
	}
	return drawbarPositions
}

func FormatDrawbarPositions(positions []int) string {
	return fmt.Sprintf(
		"%d%d %d%d%d%d %d%d%d",
		positions[0],
		positions[1],
		positions[2],
		positions[3],
		positions[4],
		positions[5],
		positions[6],
		positions[7],
		positions[8],
	)
}
