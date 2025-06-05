/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/genrpg/system/mothership"
	"github.com/genrpg/system/pirateborg"
	"github.com/genrpg/system/shadowdark"
	"github.com/spf13/cobra"
)

var extra bool
var optimized bool

// charCmd represents the char command
var charCmd = &cobra.Command{
	Use:   "char",
	Short: "Generate a random character.",
	Long:  `Generate a random character from a variety of systems. (Only Pirate Borg atm)`,
	Run: func(cmd *cobra.Command, args []string) {
		resolveSystem()
		switch system {
		case pirateBorgName:
			pirateborg.GenerateCharacter(extra)
		case shadowDarkName:
			shadowdark.GenerateCharacter(optimized, extra)
		case mothershipName:
			mothership.NewPC()
		default:
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(charCmd)
	charCmd.Flags().BoolVarP(&extra, "extra", "e", false, "Generates a character from extra character options.")
	// charCmd.Flags().BoolVarP(&optimized, "optimized", "op", false, "Generates a character picking optimal options depending on stats.")
}
