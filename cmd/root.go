/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"mittere/requester"
	"os"

	"github.com/spf13/cobra"
)

var (
	urlPath       string
	htppMethod    string
	filepath      string
	selectedColor string
	colorize      bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "mittere",
	Short:   "",
	Long:    ``,
	Version: "0.1",
	Run: func(cmd *cobra.Command, args []string) {
		requester.ExecRequest(htppMethod, urlPath,
			filepath, selectedColor, colorize)
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mittere.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false,
		"Help message for toggle")

	//url
	rootCmd.Flags().StringVarP(&urlPath, "url", "u", "",
		"Url to request data from")
	//method
	rootCmd.Flags().StringVarP(&htppMethod, "method", "m",
		"", "HTTP method")

	//filepath
	rootCmd.Flags().StringVarP(&filepath, "filepath", "f",
		"", "Path to file containing data for the request")
	rootCmd.MarkFlagRequired("filepath")
	//TODO: add custom error
	// if err := rootCmd.MarkFlagRequired("url"); err != nil {

	//colorize
	rootCmd.Flags().BoolVarP(&colorize, "colorize", "c", false,
		"Colorizes the output")

	//color
	rootCmd.Flags().StringVarP(&selectedColor, "color", "C",
		"",
		"Selected color. Values = red, green, blue, yellow, white and black.")

}
