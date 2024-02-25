package knave

import (
	"fmt"
	"math/rand"
	"strings"
)

var spellFormulaeOutput = []string{
	"[Element] [Form]",
	"[Effect] [Form]",
	"[Effect] [Element]",
	"The [Quality] [Element] [Form]",
	"The [Quality] [Effect] [Form]",
	"The [Quality] [Effect] [Element]",
	"[Wizard name]’s [Element] [Form]",
	"[Wizard name]’s [Effect] [Form]",
	"[Wizard name]’s [Effect] [Element]",
	"[Wizard name]’s [Quality] [Element] [Form]",
	"[Wizard name]’s [Quality] [Effect] [Form]",
	"[Wizard name]’s [Quality] [Effect] [Element]",
}

func getSpell(i int, spellParts ...string) string {
	switch {
	case i <= 2:
		return fmt.Sprintf("%s %s", spellParts[0], spellParts[1])
	case i <= 5:
		return fmt.Sprintf("The %s %s %s", spellParts[0], spellParts[1], spellParts[2])
	case i <= 8:
		return fmt.Sprintf("%s’s %s %s", spellParts[0], spellParts[1], spellParts[2])
	default:
		return fmt.Sprintf("%s’s %s %s %s", spellParts[0], spellParts[1], spellParts[2], spellParts[3])
	}
}

var spellFormulae = [][]func(int) string{
	{getElement, getForm},
	{getEffect, getForm},
	{getEffect, getElement},
	{getQuality, getElement, getForm},
	{getQuality, getEffect, getForm},
	{getQuality, getEffect, getElement},
	{getWizardName, getElement, getForm},
	{getWizardName, getEffect, getForm},
	{getWizardName, getEffect, getElement},
	{getWizardName, getQuality, getElement, getForm},
	{getWizardName, getQuality, getEffect, getForm},
	{getWizardName, getQuality, getEffect, getElement},
}

func getElement(i int) string {
	return elements[i]
}

func getForm(i int) string {
	return forms[i]
}

func getEffect(i int) string {
	return effects[i]
}

func getQuality(i int) string {
	return qualities[i]
}

func getWizardName(i int) string {
	return wizardNames[i]
}

// GetSpell returns a random spell
func GetRandomSpell(verbose bool) string {
	verboseOutput := ""
	spellParts := []string{}
	initialRoll := rand.Intn(12)
	spellFormulaOutput := spellFormulaeOutput[initialRoll]
	verboseOutput += fmt.Sprintf("Roll: %d\nFormula: %s\n\n", initialRoll+1, spellFormulaOutput)
	liminalString := strings.ReplaceAll(spellFormulaOutput, "] [", " ")
	liminalString = strings.ReplaceAll(liminalString, "]’s [", " ")
	liminalString = strings.ReplaceAll(liminalString, "The ", "")
	liminalString = strings.ReplaceAll(liminalString, "]", "")
	liminalString = strings.ReplaceAll(liminalString, "[", "")
	spellFormulaParts := strings.Split(liminalString, " ")
	for i, f := range spellFormulae[initialRoll] {
		roll := rand.Intn(100)
		spellPart := f(roll)
		verboseOutput += fmt.Sprintf("Table: %s | Roll: %d | Result: %s\n\n", spellFormulaParts[i], roll+1, spellPart)
		spellParts = append(spellParts, spellPart)
	}
	spell := getSpell(initialRoll, spellParts...)
	if verbose {
		verboseOutput += fmt.Sprintf("Spell: %s", spell)
		return verboseOutput
	}
	return spell
}

func GetSpell(i int64, verbose bool) string {
	verboseOutput := ""
	intelligence := int(i)
	roll := rand.Intn(100)
	spell := setSpells[roll]
	if intelligence > 0 {
		spell = strings.ReplaceAll(spell, "INT", fmt.Sprintf("%d", intelligence))
	}
	if verbose {
		verboseOutput += fmt.Sprintf("Roll:%d | Spell: %s", roll, spell)
		return verboseOutput
	}
	return spell
}
