package mazerats

import (
	"fmt"
	"strings"

	"github.com/genrpg/utils"
	"github.com/ttacon/chalk"
)

func getPhysicalEffect(i, ii int) string {
	return physicalEffect[i*6+ii]
}

func getPhysicalElement(i, ii int) string {
	return physicalElement[i*6+ii]
}

func getPhysicalForm(i, ii int) string {
	return physicalForm[i*6+ii]
}

func getEtherealEffect(i, ii int) string {
	return etherealEffect[i*6+ii]
}

func getEtherealElement(i, ii int) string {
	return etherealElement[i*6+ii]
}

func getEtherealForm(i, ii int) string {
	return etherealForm[i*6+ii]
}

var spellRecipe = [][2]func(int, int) string{
	{getPhysicalEffect, getPhysicalForm},
	{getPhysicalEffect, getPhysicalForm},
	{getEtherealEffect, getPhysicalForm},
	{getEtherealEffect, getEtherealForm},
	{getPhysicalElement, getPhysicalForm},
	{getPhysicalElement, getEtherealForm},
	{getEtherealElement, getPhysicalForm},
	{getEtherealElement, getEtherealForm},
	{getPhysicalEffect, getPhysicalElement},
	{getPhysicalEffect, getEtherealElement},
	{getEtherealEffect, getPhysicalElement},
	{getEtherealEffect, getEtherealElement},
}

func getSpellRecipe(i, ii int) [2]func(int, int) string {
	return spellRecipe[i/3*6+ii]
}
func getSpellRecipeOutput(i, ii int) string {
	return spellRecipeOutput[i/3*6+ii]
}

// GetSpell returns a random spell
func GetRandomSpell(verbose bool) string {
	verboseOutput := ""
	spellParts := []string{}
	rolls := [2]int{}
	for i := 0; i < 2; i++ {
		rolls[i] = utils.TableDie(6)
	}
	spellRecipeOutput := getSpellRecipeOutput(rolls[0], rolls[1])
	spellRecipeOutputParts := strings.Split(spellRecipeOutput, " + ")
	verboseOutput += fmt.Sprintf("Rolls: %d, %d\nRecipe: %s\n\n", rolls[0]+1, rolls[1]+1, spellRecipeOutput)
	// verboseOutput += "Table | Rolls | Result\n"
	for i, f := range getSpellRecipe(rolls[0], rolls[1]) {
		for ii := 0; ii < 2; ii++ {
			rolls[ii] = utils.TableDie(6)
		}
		spellParts = append(spellParts, f(rolls[0], rolls[1]))
		verboseOutput += fmt.Sprintf("%s | %s | %s\n\n",
			chalk.Bold.TextStyle(spellRecipeOutputParts[i]),
			chalk.Bold.TextStyle(fmt.Sprintf("[%d, %d]", rolls[0]+1, rolls[1]+1)),
			chalk.Bold.TextStyle(spellParts[i]))
	}
	spell := utils.SpellStyle(strings.Join(spellParts, (" ")))
	if verbose {
		verboseOutput += fmt.Sprintf("Spell: %s", spell)
		return verboseOutput
	}
	return spell
}
