/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/genrpg/dungeon"
	"github.com/spf13/cobra"
)

// mapCmd represents the stock command
var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "Make a map",
	Long:  `Generate a dungeon map`,
	Run: func(cmd *cobra.Command, args []string) {
		dungeon.Generate()
	},
}

func init() {
	rootCmd.AddCommand(mapCmd)

}
