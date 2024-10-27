/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/genrpg/system/pirateborg"
	"github.com/spf13/cobra"
)

// IslandCmd represents the island command
var IslandCmd = &cobra.Command{
	Use:   "island",
	Short: "Generate a random island.",
	Long:  `Generate a random island from a variety of systems. (Only Pirate Borg atm)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pirateborg.GetIsland())
	},
}
