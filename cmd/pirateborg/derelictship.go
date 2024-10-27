/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/genrpg/system/pirateborg"
	"github.com/spf13/cobra"
)

// derelictShipCmd represents the derelict ship command
var DerelictShipCmd = &cobra.Command{
	Use:   "derelict",
	Short: "Generate a random derelict ship.",
	Long:  `Generate a random derelict ship from a variety of systems. (Only Pirate Borg atm)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pirateborg.GetDerelictShip())
	},
}
