package mazerats

import (
	"fmt"
	"math/rand"
	"strings"
)

var physicalEffect = []string{
	"Animating", "Attracting", "Binding", "Blossoming", "Consuming", "Creeping", "Crushing", "Diminishing",
	"Dividing", "Duplicating", "Enveloping", "Expanding", "Fusing", "Grasping", "Hastening", "Hindering",
	"Illuminating", "Imprisoning", "Levitating", "Opening", "Petrifying", "Phasing", "Piercing", "Pursuing",
	"Reflecting", "Regenerating", "Rending", "Repelling", "Resurrecting", "Screaming", "Sealing", "Shapeshifting",
	"Shielding", "Spawning", "Transmuting", "Transporting",
}

var physicalElement = []string{
	"Acid", "Amber", "Bark", "Blood", "Bone", "Brine", "Clay", "Crow", "Crystal", "Ember", "Flesh", "Fungus", "Glass",
	"Honey", "Ice", "Insect", "Wood", "Lava", "Moss", "Obsidian", "Oil", "Poison", "Rat", "Salt", "Sand", "Sap",
	"Serpent", "Slime", "Stone", "Tar", "Thorn", "Vine", "Water", "Wine", "Wood", "Worm",
}

var physicalForm = []string{
	"Altar", "Armor", "Arrow", "Beast", "Blade", "Cauldron", "Chain", "Chariot", "Claw", "Cloak", "Colossus", "Crown",
	"Elemental", "Eye", "Fountain", "Gate", "Golem", "Hammer", "Horn", "Key", "Mask", "Monolith", "Pit", "Prison",
	"Sentinel", "Servant", "Shield", "Spear", "Steed", "Swarm", "Tentacle", "Throne", "Torch", "Trap", "Wall", "Web",
}

var etherealEffect = []string{
	"Avenging", "Banishing", "Bewildering", "Blinding", "Charming", "Communicating", "Compelling", "Concealing",
	"Deafening", "Deceiving", "Deciphering", "Disguising", "Dispelling", "Emboldening", "Encoding", "Energizing",
	"Enlightening", "Enraging", "Excruciating", "Foreseeing", "Intoxicating", "Maddening", "Mesmerizing",
	"Mindreading", "Nullifying", "Paralyzing", "Revealing", "Revolting", "Scrying", "Silencing", "Soothing",
	"Summoning", "Terrifying", "Warding", "Wearying", "Withering",
}

var etherealElement = []string{
	"Ash", "Chaos", "Distortion", "Dream", "Dust", "Echo", "Ectoplasm", "Fire", "Fog", "Ghost", "Harmony", "Heat",
	"Light", "Lightning", "Memory", "Mind", "Mutation", "Negation", "Plague", "Plasma", "Probability", "Rain", "Rot",
	"Shadow", "Smoke", "Snow", "Soul", "Star", "Stasis", "Steam", "Thunder", "Time", "Void", "Warp", "Whisper", "Wind",
}

var etherealForm = []string{
	"Aura", "Beacon", "Beam", "Blast", "Blob", "Bolt", "Bubble", "Call", "Cascade", "Circle", "Cloud", "Coil", "Cone",
	"Cube", "Dance", "Disk", "Field", "Form", "Gaze", "Loop", "Moment", "Nexus", "Portal", "Pulse", "Pyramid", "Ray",
	"Shard", "Sphere", "Spray", "Storm", "Swarm", "Torrent", "Touch", "Vortex", "Wave", "Word",
}

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

var spellRecipeOutput = []string{
	"Physical Effect + Physical Form",
	"Physical Effect + Physical Form",
	"Ethereal Effect + Physical Form",
	"Ethereal Effect + Ethereal Form",
	"Physical Element + Physical Form",
	"Physical Element + Ethereal Form",
	"Ethereal Element + Physical Form",
	"Ethereal Element + Ethereal Form",
	"Physical Effect + Physical Element",
	"Physical Effect + Ethereal Element",
	"Ethereal Effect + Physical Element",
	"Ethereal Effect + Ethereal Element",
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
		rolls[i] = rand.Intn(6)
	}
	spellRecipeOutput := getSpellRecipeOutput(rolls[0], rolls[1])
	spellRecipeOutputParts := strings.Split(spellRecipeOutput, " + ")
	verboseOutput += fmt.Sprintf("Rolls: %d, %d\nRecipe: %s\n\n", rolls[0]+1, rolls[1]+1, spellRecipeOutput)
	for i, f := range getSpellRecipe(rolls[0], rolls[1]) {
		for ii := 0; ii < 2; ii++ {
			rolls[ii] = rand.Intn(6)
		}
		verboseOutput += fmt.Sprintf("%s\nRolls: %d, %d\n", spellRecipeOutputParts[i], rolls[0]+1, rolls[1]+1)
		spellParts = append(spellParts, f(rolls[0], rolls[1]))
		verboseOutput += fmt.Sprintf("Table result: %s\n\n", spellParts[i])
	}
	spell := strings.Join(spellParts, (" "))
	if verbose {
		verboseOutput += fmt.Sprintf("Spell: %s", spell)
		return verboseOutput
	}
	return spell
}
