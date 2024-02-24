/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	mazerats "github.com/genrpg/system/maze_rats"
	"github.com/spf13/cobra"
)

// spellCmd represents the spell command
var spellCmd = &cobra.Command{
	Use:   "spell",
	Short: "Generate a random spell.",
	Long:  `Generate a random spell from a variety of systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		var spell string
		switch system {
		case "mr":
			spell = mazerats.GetRandomSpell(verbose)
		case "mazerats":
			spell = mazerats.GetRandomSpell(verbose)
		}
		fmt.Printf("\nUsing %s\n%s\n", systemMap[system], spell)
	},
}

func init() {
	rootCmd.AddCommand(spellCmd)
}
