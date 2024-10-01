/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/genrpg/system/knave"
	mazerats "github.com/genrpg/system/maze_rats"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var spellRandom bool
var spellInt int64

// spellCmd represents the spell command
var spellCmd = &cobra.Command{
	Use:   "spell",
	Short: "Generate a random spell.",
	Long:  `Generate a random spell from a variety of systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		var spell string
		switch {
		case system == "mr" || system == "mazerats":
			spell = mazerats.GetRandomSpell(verbose)
		case system == "knave" || system == "k":
			if spellRandom {
				spell = knave.GetRandomSpell(verbose)
			} else {
				spell = knave.GetSpell(spellInt, verbose)
			}
		}
		if verbose {
			fmt.Printf("\n%s\n\n%s\n", chalk.Underline.TextStyle(fmt.Sprintf("Using %s", systemMap[system])), spell)
		} else {
			fmt.Println(spell)
		}
	},
}

func init() {
	rootCmd.AddCommand(spellCmd)
	spellCmd.Flags().BoolVarP(&spellRandom, "random", "r", false, "Generates a random spell rather than a predetermined one. Knave only.")
	spellCmd.Flags().Int64VarP(&spellInt, "intelligence", "i", 0, "Set the intelligence score to get already calculated values. Knave only.")
}
