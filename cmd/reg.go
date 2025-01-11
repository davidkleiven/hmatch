/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/davidkleiven/hmatch/pkg/standard"
	"github.com/spf13/cobra"
)

// regCmd represents the reg command
var regCmd = &cobra.Command{
	Use:   "reg",
	Short: "List suggested registrations that matches pipe on the pipe organ",
	Long: `List registrations from https://b3world.com/hammond-drawbar-settings.html
	that matches pipes from the pipe organ. If not pattern is given, all registrations
	are listed

	Example:
	hmatch reg tib

	will list all hammond organ drawbar positions that closely resembles the sound
	of various pipes that contains the substring "tib"

	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var registrations standard.Registrations
		if len(args) != 1 {
			registrations = standard.AllRegistations()
		} else {
			registrations = standard.FindRegistration(args[0])
		}

		fmt.Printf("%s\n", registrations.ToString())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(regCmd)
}
