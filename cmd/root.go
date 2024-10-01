/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var system string
var verbose bool

var mazeRatsName = "Maze Rats by Ben Milton"
var knaveName = "Knave 2nd edition by Ben Milton"
var shadowDarkName = "ShadowDark by Kelsey Dionne"

var systemMap = map[string]string{
	"mr":         mazeRatsName,
	"mazerats":   mazeRatsName,
	"knave":      knaveName,
	"k":          knaveName,
	"sd":         shadowDarkName,
	"shadowdark": shadowDarkName,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "genrpg",
	Short: "A generator tool for RPG games.",
	Long:  `This is a tool for generating random spells, items, and other things for RPG games.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	rootCmd.PersistentFlags().StringVarP(&system, "system", "s", "k", "System to use (mr, knave, sd, etc.)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Explains all of the tables used and numbers rolled to generate the output.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
