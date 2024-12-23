/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/davidkleiven/hmatch/pkg"
	"github.com/spf13/cobra"
)

// combineCmd represents the combine command
var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Combines multiple drawbar settings",
	Long: `Combine multiple drawbar settings into one. Drawbar settings are combined
	based on the criteria that the power of the final setting should be equal to the
	power of the individual

	Example:
	hmatch combine 008000123,010000000
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("combination is missing")
		}

		combinations, err := pkg.ParseDrawbarPositions(args[0])
		if err != nil {
			return err
		}

		summedPowerCombination := pkg.Combine(combinations)
		fmt.Printf("%s\n", pkg.FormatDrawbarPositions(summedPowerCombination[:]))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(combineCmd)
}
