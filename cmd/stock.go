/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/genrpg/system/ose"
	"github.com/genrpg/system/shadowdark"
	"github.com/spf13/cobra"
)

var rooms int64

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "Stock a dungeon.",
	Long:  `Stock a dungeon with monsters, traps, and treasure.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch rooms {
		case 0:
			fmt.Println("Nothing to stock.")
		case 1000:
			fmt.Println("Too many rooms to stock, will default to 999.")
			rooms = 999
			fallthrough
		default:
			switch {
			case system == "sd" || system == "shadowdark":
				shadowdark.Stocking(rooms)
			default:
				system = "ose"
				ose.Stocking(rooms)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(stockCmd)

	stockCmd.Flags().Int64VarP(&rooms, "number of rooms", "n", 1, "Number of rooms to stock.")
}
