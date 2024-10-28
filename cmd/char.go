/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/genrpg/system/pirateborg"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var extra bool

// charCmd represents the char command
var charCmd = &cobra.Command{
	Use:   "char",
	Short: "Generate a random character.",
	Long:  `Generate a random character from a variety of systems. (Only Pirate Borg atm)`,
	Run: func(cmd *cobra.Command, args []string) {
		pc := pirateborg.NewCharacter(extra)
		fmt.Printf("%s %s\n%s %s\n%s\n%s\n",
			pirateborg.Red("Name:"), pc.Name,
			pirateborg.Red("Class:"), pc.Class,
			pc.Class.Description(),
			pirateborg.Red("Features:"))
		for _, feature := range pc.Features {
			fmt.Println(feature)
		}
		fmt.Printf("\n%s %d\n", pirateborg.Red("HP:"), pc.HP)
		fmt.Printf("%s\n%s\n", pirateborg.Red("Stats:"), pc.Stats)
		fmt.Printf("%s %s\n", pirateborg.Red("Devil's Luck:"), pc.Class.GetDevilsLuck())

		fmt.Printf("\n%s\n%s", pirateborg.Red("Weapons:"), pc.Weapons)
		fmt.Printf("%s %s\n", pirateborg.Red("Clothing:"), pc.Clothing)
		fmt.Printf("%s %s\n", pirateborg.Red("Hat:"), pc.Hat)
		fmt.Printf("%s\n%s", pirateborg.Red("Equipment:"), pc.Gear)
		fmt.Printf("%s %s\n", pirateborg.Red("Money:"), pc.Money)

		fmt.Printf("\n%s\n", pc.Character)
		fmt.Printf("%s %s\n", chalk.Yellow.Color("Thing of Importance:"), pc.ThingOfImportance)
	},
}

func init() {
	rootCmd.AddCommand(charCmd)
	charCmd.Flags().BoolVarP(&extra, "extra", "e", false, "Generates a character from extra character options. Pirate Borg only.")
}
