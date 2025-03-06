/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/genrpg/system/shadowdark"
	"github.com/spf13/cobra"
)

// monsterCmd represents the monster command
var monsterCmd = &cobra.Command{
	Use:   "monster",
	Short: "Get a monster",
	Long:  `Get or generate a monster.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) == 0 {
			err = fmt.Errorf("Input a number")
			panic(err)
		}
		level, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}
		output := shadowdark.MonsterGenerator(int(level))
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(monsterCmd)
}
