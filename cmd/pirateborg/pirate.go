/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/genrpg/system/pirateborg"
	"github.com/spf13/cobra"
)

// VesselCmd represents the vessel command
var PirateCmd = &cobra.Command{
	Use:   "pirate",
	Short: "Generate a random pirate.",
	Long:  `Generate a random pirate from a variety of systems. (Only Pirate Borg atm)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pirateborg.GeneratePirate())
	},
}
