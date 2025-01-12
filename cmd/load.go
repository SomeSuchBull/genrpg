/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/genrpg/system/shadowdark"
	"github.com/spf13/cobra"
)

// spellCmd represents the spell command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load data for something",
	Long:  `Generate code for further development, usually some content.`,
	Run: func(cmd *cobra.Command, args []string) {
		outputFile, err := os.Create("data/monsters.go")
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()
		shadowdark.LoadMonsters(outputFile)
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
