package knave

import (
	"fmt"
	"math/rand"
)

// TIP: Roll twice on these tables and combine the results.
var engine = map[int]func() string{
	0: getStockingTreasure, 1: getStockingTreasure, 2: getStockingMonster,
	3: getStockingMonster, 4: getSpecial, 5: getTrap}

func Stocking(rooms int64) {
	for i := int64(0); i < rooms; i++ {
		roomContents := ""
		roll := rand.Intn(6)
		f := engine[roll]
		roomContents += f()
		fmt.Printf("%03d: %s\n", i+1, roomContents)
	}
}

func getStockingTreasure() string {
	return "Treasure"
}

func getStockingMonster() string {
	return "Monster"
}

func getSpecial() string {
	return "Special"
}

func getTrap() string {
	return "Trap"
}

var recurs func(string, *string) string

var getRandomSpellProxy func(bool) string

func init() {
	recurs = recursiveTableRoll
	getRandomSpellProxy = GetRandomSpell
}

func recursiveTableRoll(key string, verboseOutput *string) string {
	if referencedTables[key] == nil {
		return key
	}
	table := key
	roll := rand.Intn(100)
	tableResults := referencedTables[key]
	var result string
	if tableResults["function"] != nil {
		result = tableResults["function"].(func(int) string)(roll)
	} else {
		table := tableResults["table"].([]string)
		result = table[roll]
	}
	if verboseOutput != nil {
		*verboseOutput += tableRoll(table, roll, result)
	}
	format := "%s"
	if tableResults["format"] != nil {
		format = tableResults["format"].(string)
	}
	return fmt.Sprintf(format, recursiveTableRoll(result, verboseOutput))
}

func tableRoll(table string, roll int, result string) string {
	return fmt.Sprintf("-----\nTable:  %s\nRoll:   %-3.02d\nResult: %s\n\n", table, roll+1, result)
}

func getSpellbook(i int) string {
	// return spellbooks[i]
	return "A spellbook, todo"
}
func getSpellScroll(i int) string {
	return "Scroll of " + setSpells[i]
}

func getInn(i int) string {
	innNamePart1 := recurs(innName1[i], nil)
	innNamePart2 := recurs(innName2[rand.Intn(len(innName2))], nil)
	return innNamePart1 + " " + innNamePart2
}

func getName(i int) string {
	if rand.Intn(2) == 0 {
		return maleNames[i]
	}
	return femaleNames[i]
}

func getSurname(i int) string {
	return surname1[i] + surname2[rand.Intn(len(surname2))]
}

func getRandSpell(i int) string {
	return getRandomSpellProxy(false)
}

var referencedTables = map[string]map[string]any{
	"Travel shift (p. 9)":   {"table": travelShifts},
	"Sign (p. 10)":          {"table": signs},
	"Location (p. 10)":      {"table": locations},
	"Place trait (p. 11)":   {"table": placeTraits},
	"Structure (p. 11)":     {"table": structures},
	"Delve shift (p. 14)":   {"table": delveShifts},
	"Room (p. 14)":          {"table": rooms},
	"Room detail (p. 15)":   {"table": roomDetails},
	"Room theme (p. 15)":    {"table": roomThemes},
	"Trap effect (p. 16)":   {"table": trapEffects},
	"Dungeon (p. 16)":       {"table": dungeons},
	"Hazard (p. 17)":        {"table": hazards},
	"Mechanism (p. 17)":     {"table": mechanisms},
	"Spell scroll (p. 22)":  {"function": getSpellScroll},
	"Spellbook (p. 22)":     {"function": getSpellbook},
	"Spell (pp. 22-25)":     {"function": getRandSpell},
	"Wiz. name (p. 27)":     {"table": wizardNames},
	"Effect (p. 28)":        {"table": effects},
	"Effect rain (p. 28)":   {"table": effects, "format": "%s rain"},
	"Effect aura (p. 28)":   {"table": effects, "format": "%s aura"},
	"Effect blast (p. 28)":  {"table": effects, "format": "%s blast"},
	"Effect bolt (p. 28)":   {"table": effects, "format": "%s bolt"},
	"Effect ray (p. 28)":    {"table": effects, "format": "%s ray"},
	"Quality (p. 28)":       {"table": qualities},
	"Element (p. 29)":       {"table": elements},
	"Element rain (p. 29)":  {"table": elements, "format": "%s rain"},
	"Element field (p. 29)": {"table": elements, "format": "%s field"},
	"Element flow (p. 29)":  {"table": elements, "format": "%s flow"},
	"Element skin (p. 29)":  {"table": elements, "format": "%s skin"},
	"Element trail (p. 29)": {"table": elements, "format": "%s trail"},
	"Elem. form (p. 29)":    {"table": elements, "format": "%s form"},
	"Elem. blood (p. 29)":   {"table": elements, "format": "%s blood"},
	"Elem. body (p. 29)":    {"table": elements, "format": "%s body"},
	"Elem. breath (p. 29)":  {"table": elements, "format": "%s breath"},
	"Element aura (p. 29)":  {"table": elements, "format": "%s aura"},
	"Element blast (p. 29)": {"table": elements, "format": "%s blast"},
	"Element bolt (p. 29)":  {"table": elements, "format": "%s bolt"},
	"Elem. control (p. 29)": {"table": elements, "format": "%s control"},
	"Element ray (p. 29)":   {"table": elements, "format": "%s ray"},
	"Element wall (p. 29)":  {"table": elements, "format": "%s wall"},
	"Delusion (p. 30)":      {"table": delusions},
	"Mutation (p. 30)":      {"table": mutations},
	"Disaster (p. 31)":      {"table": disasters},
	"Magic school (p. 31)":  {"table": magicSchools},
	"Domain (p. 33)":        {"table": domains},
	"Fears dom. (p. 33)":    {"table": domains, "format": "Fears %s"},
	"Symbol (p. 33)":        {"table": symbols},
	"Fears symbol (p. 33)":  {"table": symbols, "format": "Fears %s"},
	"Potion (p. 35)":        {"table": potions},
	"Potion recipe (p. 35)": {"table": potions, "format": "Recipe for a potion of %s"},
	"Taste (p. 36)":         {"table": tastes},
	"Taste rain (p. 36)":    {"table": tastes, "format": "%s rain"},
	"Texture (p. 36)":       {"table": textures},
	"Texture body (p. 36)":  {"table": textures, "format": "%s body"},
	"Texture rain (p. 36)":  {"table": textures, "format": "%s rain"},
	"Skin texture (p. 36)":  {"table": textures, "format": "Skin %s"},
	"Texture trail (p. 36)": {"table": textures, "format": "%s trail"},
	"Ingredient (p. 37)":    {"table": ingredients},
	"Color (p. 37)":         {"table": colors},
	"Color skin (p. 37)":    {"table": colors, "format": "%s skin"},
	"Color sky (p. 37)":     {"table": colors, "format": "%s sky"},
	"Tool (p. 39)":          {"table": tools},
	"Misc. item (p. 39)":    {"table": miscellaneousItems},
	"Book (p. 40)":          {"table": books},
	"Book expert (p. 40)":   {"table": books, "format": "%s expert"},
	"Clothing (p. 40)":      {"table": clothing},
	"Fabric (p. 41)":        {"table": fabrics},
	"Fabric scrap (p. 41)":  {"table": fabrics, "format": "%s scrap"},
	"Material (p. 42)":      {"table": materials},
	"Treasure (p. 42)":      {"table": treasures},
	"Weapon (p. 43)":        {"table": weapons},
	"Weapon rain (p. 43)":   {"table": weapons, "format": "%s rain"},
	"Item trait (p. 43)":    {"table": itemTraits},
	"City event (p. 46)":    {"table": cityEvents},
	"City theme (p. 46)":    {"table": cityThemes},
	"Building (p. 47)":      {"table": buildings},
	"Street detail (p. 47)": {"table": streetDetails},
	"Inn (p. 48)":           {"function": getInn},
	"Food (p. 49)":          {"table": food},
	"Food addict (p. 49)":   {"table": food, "format": "%s addict"},
	"Food rain (p. 49)":     {"table": food, "format": "%s rain"},
	"Food scraps (p. 49)":   {"table": food, "format": "%s scraps"},
	"Faction (p. 50)":       {"table": factions},
	"Faction ally (p. 50)":  {"table": factions, "format": "%s ally"},
	"Faction trait (p. 50)": {"table": factions, "format": "%s trait"},
	"Faction war (p. 50)":   {"table": factions, "format": "%s war"},
	"Mission (p. 51)":       {"table": missions},
	"Reward (p. 51)":        {"table": rewards},
	"Archetype (p. 53)":     {"table": archetypes},
	"Name (pp. 54-55)":      {"function": getName},
	"Surname (p. 55)":       {"function": getSurname},
	"NPC detail (p. 56)":    {"table": npcDetails},
	"Personality (p. 56)":   {"table": personalities},
	"Goal (p. 57)":          {"table": goals},
	"Profession (p. 57)":    {"table": professions},
	"Asset (p. 58)":         {"table": assets},
	"Liability (p. 58)":     {"table": liabilities},
	"Mannerism (p. 59)":     {"table": mannerisms},
	"Monster (p. 61)":       {"table": monsters},
	"Monster ally (p. 61)":  {"table": monsters, "format": "%s ally"},
	"Monster skin (p. 61)":  {"table": monsters, "format": "%s skin"},
	"Animal (p. 64)":        {"table": animals},
	"Animals (p. 64)":       {"table": animals},
	"Animal body (p. 64)":   {"table": animals, "format": "%s body"},
	"Animal ears (p. 64)":   {"table": animals, "format": "%s ears"},
	"Animal eyes (p. 64)":   {"table": animals, "format": "%s eyes"},
	"Animal form (p. 64)":   {"table": animals, "format": "%s form"},
	"Animal head (p. 64)":   {"table": animals, "format": "%s head"},
	"Animal limb (p. 64)":   {"table": animals, "format": "%s limb"},
	"Animal scent (p. 64)":  {"table": animals, "format": "%s scent"},
	"Animal skin (p. 64)":   {"table": animals, "format": "%s skin"},
	"Animal tail (p. 64)":   {"table": animals, "format": "%s tail"},
	"Animal teeth (p. 64)":  {"table": animals, "format": "%s teeth"},
	"Animal voice (p. 64)":  {"table": animals, "format": "%s voice"},
	"Anim. form (p. 64)":    {"table": animals, "format": "%s form"},
	"Giant anim. (p. 64)":   {"table": animals, "format": "Giant %s"},
	"Organ (p. 64)":         {"table": organs},
	"Monster trait (p. 65)": {"table": monsterTraits},
	"Mon. trait (p. 65)":    {"table": monsterTraits},
	"Power (p. 65)":         {"table": powers},
	"Scent (p. 66)":         {"table": scents},
	"Sound (p. 66)":         {"table": sounds},
	"Tactic (p. 67)":        {"table": tactics},
	"Weakness (p. 67)":      {"table": weaknesses},
}
