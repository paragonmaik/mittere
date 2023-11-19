/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mittere/requester"
	"os"

	"github.com/spf13/cobra"
)

var (
	urlPath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "mittere",
	Short:   "",
	Long:    ``,
	Version: "0.1",
	Run: func(cmd *cobra.Command, args []string) {
		requester.Request(urlPath)
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

	rootCmd.Flags().StringVarP(&urlPath, "url", "u", "",
		"Url to request data from")

	if err := rootCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
		//TODO: add custom error
	}
}
