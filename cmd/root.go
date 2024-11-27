/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hmatch",
	Short: "Matches hammond drawbars to sound recording",
	Long: `hmatch is an application that calculates best hammond drawbar settings
	that matches the frequencey spectrum of a recording. The method is based on
	the article: http://www.stefanv.com/electronics/hammond_drawbar_science.html

	Example:
	hmatch fit audio.wav`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
