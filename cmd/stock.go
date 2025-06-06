/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/genrpg/system/knave"
	"github.com/genrpg/system/shadowdark"
	"github.com/spf13/cobra"
)

var rooms int64
var level int64

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "Stock a dungeon.",
	Long:  `Stock a dungeon with monsters, traps, treasure, and other.`,
	Run: func(cmd *cobra.Command, args []string) {
		resolveSystem()
		switch {
		case rooms == 0:
			fmt.Println("Nothing to stock.")
		case rooms > 1000:
			fmt.Println("Too many rooms to stock, will default to 999.")
			rooms = 999
			fallthrough
		default:
			switch system {
			case shadowDarkName:
				shadowdark.Stocking(rooms, int(level))
			default:
				system = "knave"
				resolveSystem()
				knave.Stocking(rooms, level, verbose)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(stockCmd)

	stockCmd.Flags().Int64VarP(&rooms, "number of rooms", "n", 1, "Number of rooms to stock.")
	stockCmd.Flags().Int64VarP(&level, "level of dungeon", "l", 0, "Level of dungeon to stock.")
}
