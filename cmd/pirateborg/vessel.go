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
var VesselCmd = &cobra.Command{
	Use:   "vessel",
	Short: "Generate a random vessel.",
	Long:  `Generate a random vessel from a variety of systems. (Only Pirate Borg atm)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pirateborg.GetVessel())
	},
}
