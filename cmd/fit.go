/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cmp"
	"errors"
	"log"
	"slices"

	"github.com/davidkleiven/hmatch/pkg"
	"github.com/spf13/cobra"
)

// fitCmd represents the fit command
var fitCmd = &cobra.Command{
	Use:   "fit <audiofile>",
	Short: "Finds best drawbar registration to a recording",
	Long: `Finds best drawbar registration to a recording

	Example:
	hmatch fit audio.wav
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("require audio file")
		}

		fname := args[0]

		hannWindow, err := pkg.IntFromFlag(cmd.Flag("windowLength"))

		if err != nil {
			return err
		}

		data, err := pkg.ReadFromFile(fname)
		if err != nil {
			return err
		}
		peaks := pkg.DetectSpectrumPeaks(data.Mean(), data.SampleRate, hannWindow)
		log.Printf("%s\n", pkg.PeakToString(peaks))

		drawbarFeet := pkg.DrawbarFeets()
		results := []pkg.DrawbarPositionResult{}
		for _, maxPeakDrawBar := range []int{0, 1, 2, 3, 4} {
			positions := pkg.FindRegistration(peaks, maxPeakDrawBar)
			results = append(results, positions)
			log.Printf(
				"Closest drawbar position when assigning peak to %.2f' (max. freq. dev %.2f Hz):\n%s\n",
				drawbarFeet[maxPeakDrawBar],
				positions.MaxFreqDev,
				pkg.FormatDrawbarPositions(positions.Positions),
			)
		}

		bestResult := slices.MinFunc(results, func(a pkg.DrawbarPositionResult, b pkg.DrawbarPositionResult) int {
			return cmp.Compare(a.MaxFreqDev, b.MaxFreqDev)
		})

		log.Printf(
			"Best result was: %v (max. freq. dev: %.2f Hz)\n",
			pkg.FormatDrawbarPositions(bestResult.Positions),
			bestResult.MaxFreqDev,
		)

		return nil
	},
}

func init() {
	fitCmd.PersistentFlags().
		Int32("windowLength", 1024, "Length of the Hann window used when calculating the frequency spectrum")
	rootCmd.AddCommand(fitCmd)
}
