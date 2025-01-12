/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/genrpg/system/shadowdark"
	"github.com/spf13/cobra"
)

// showCmd represents the stock command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show something.",
	Long:  `Show a whole list of things.`,
	Run: func(cmd *cobra.Command, args []string) {
		shadowdark.ShowDistribution()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

}
