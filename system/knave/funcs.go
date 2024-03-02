package knave

import (
	"fmt"
	"math/rand"
)

var recurs func(string, *string) string

var getRandomSpellProxy func(bool) string

func init() {
	recurs = recursiveTableRoll
	getRandomSpellProxy = GetRandomSpell
}

func recursiveTableRoll(key string, verboseOutput *string) string {
	if tableFuncs[key] == nil {
		return key
	}
	table := key
	roll := rand.Intn(100)
	key = tableFuncs[key](roll)
	if verboseOutput != nil {
		*verboseOutput += tableRoll(table, roll, key)
	}
	return recursiveTableRoll(key, verboseOutput)
}

func tableRoll(table string, roll int, result string) string {
	return fmt.Sprintf("-----\nTable:  %s\nRoll:   %-3.02d\nResult: %s\n\n", table, roll+1, result)
}

var tableFuncs = map[string]func(int) string{
	"Travel shift (p. 9)": getTravelShift, "Sign (p. 10)": getSign,
	"Location (p. 10)": getLocation, "Place trait (p. 11)": getPlaceTrait,
	"Structure (p. 11)": getStructure, "Delve shift (p. 14)": getDelveShift, "Room (p. 14)": getRoom,
	"Room detail (p. 15)": getRoomDetail,
	"Room theme (p. 15)":  getRoomTheme, "Trap effect (p. 16)": getTrapEffect, "Dungeon (p. 16)": getDungeon,
	"Hazard (p. 17)": getHazard, "Mechanism (p. 17)": getMechanism, "Spell scroll (p. 22)": getSpellScroll,
	"Spellbook (p. 22)": getSpellbook, "Spell (pp. 22-25)": getRandSpell, "Wiz. name (p. 27)": getWizardName,
	"Effect (p. 28)": getEffect, "Effect rain (p. 28)": getEffectRain, "Effect aura (p. 28)": getEffectAura,
	"Effect blast (p. 28)": getEffectBlast, "Effect bolt (p. 28)": getEffectBolt, "Effect ray (p. 28)": getEffectRay,
	"Quality (p. 28)": getQuality,
	"Element (p. 29)": getElement, "Element rain (p. 29)": getElementRain, "Element field (p. 29)": getElementField,
	"Element flow (p. 29)": getElementFlow, "Element skin (p. 29)": getElementSkin, "Element trail (p. 29)": getElementTrail,
	"Elem. form (p. 29)":  getElementForm,
	"Elem. blood (p. 29)": getElementBlood, "Elem. body (p. 29)": getElementBody, "Elem. breath (p. 29)": getElementBreath,
	"Element aura (p. 29)": getElementAura, "Element blast (p. 29)": getElementBlast, "Element bolt (p. 29)": getElementBolt,
	"Elem. control (p. 29)": getElementControl, "Element ray (p. 29)": getElementRay, "Element wall (p. 29)": getElementWall,
	"Delusion (p. 30)": getDelusion, "Mutation (p. 30)": getMutation, "Disaster (p. 31)": getDisaster,
	"Magic school (p. 31)": getMagicSchool, "Domain (p. 33)": getDomain, "Fears dom. (p. 33)": getDomainFear,
	"Symbol (p. 33)": getSymbol, "Fears symbol (p. 33)": getSymbolFear, "Potion (p. 35)": getPotion,
	"Potion recipe (p. 35)": getPotionRecipe, "Taste (p. 36)": getTaste, "Taste rain (p. 36)": getTasteRain,
	"Texture (p. 36)": getTexture, "Texture body (p. 36)": getTextureBody, "Texture rain (p. 36)": getTextureRain,
	"Skin texture (p. 36)": getSkinTexture, "Texture trail (p. 36)": getTextureTrail,
	"Ingredient (p. 37)": getIngredient, "Color (p. 37)": getColor,
	"Color skin (p. 37)": getColorSkin, "Color sky (p. 37)": getColorSky, "Tool (p. 39)": getTool,
	"Misc. item (p. 39)": getMiscItem, "Book (p. 40)": getBook, "Book expert (p. 40)": getBookExpert,
	"Clothing (p. 40)": getClothing, "Fabric (p. 41)": getFabric, "Fabric scrap (p. 41)": getFabricScrap,
	"Material (p. 42)": getMaterials,
	"Treasure (p. 42)": getTreasure, "Weapon (p. 43)": getWeapon, "Weapon rain (p. 43)": getWeaponRain,
	"Item trait (p. 43)": getItemTrait, "City event (p. 46)": getCityEvent, "City theme (p. 46)": getCityTheme,
	"Building (p. 47)": getBuilding, "Street detail (p. 47)": getStreetDetail, "Inn (p. 48)": getInn, "Food (p. 49)": getFood,
	"Food addict (p. 49)": getFoodAddict, "Food rain (p. 49)": getFoodRain, "Food scraps (p. 49)": getFoodScraps,
	"Faction (p. 50)":      getFaction,
	"Faction ally (p. 50)": getFactionAlly, "Faction trait (p. 50)": getFactionTrait, "Faction war (p. 50)": getFactionWar,
	"Mission (p. 51)": getMission, "Reward (p. 51)": getReward, "Archetype (p. 53)": getArchetype,
	"Name (pp. 54-55)": getName, "Surname (p. 55)": getSurname, "NPC detail (p. 56)": getNPCDetail,
	"Personality (p. 56)": getPersonality, "Goal (p. 57)": getGoal, "Profession (p. 57)": getProfession,
	"Asset (p. 58)": getAsset, "Liability (p. 58)": getLiability, "Mannerism (p. 59)": getMannerism,
	"Monster (p. 61)": getMonster, "Monster ally (p. 61)": getMonsterAlly, "Monster skin (p. 61)": getMonsterSkin,
	"Animal (p. 64)": getAnimal, "Animals (p. 64)": getAnimal,
	"Animal body (p. 64)": getAnimalBody, "Animal ears (p. 64)": getAnimalEars,
	"Animal eyes (p. 64)": getAnimalEyes, "Animal form (p. 64)": getAnimalForm, "Animal head (p. 64)": getAnimalHead,
	"Animal limb (p. 64)": getAnimalLimb, "Animal scent (p. 64)": getAnimalScent, "Animal skin (p. 64)": getAnimalSkin,
	"Animal tail (p. 64)": getAnimalTail, "Animal teeth (p. 64)": getAnimalTeeth, "Animal voice (p. 64)": getAnimalVoice,
	"Anim. form (p. 64)": getAnimalForm, "Giant anim. (p. 64)": getGiantAnimal, "Organ (p. 64)": getOrgan,
	"Monster trait (p. 65)": getMonsterTrait, "Mon. trait (p. 65)": getMonsterTrait, "Power (p. 65)": getPower,
	"Scent (p. 66)": getScent, "Sound (p. 66)": getSound, "Tactic (p. 67)": getTactic, "Weakness (p. 67)": getWeakness,
}

func getSign(i int) string {
	return signs[i]
}
func getRoomDetail(i int) string {
	return roomDetails[i]
}
func getStreetDetail(i int) string {
	return streetDetails[i]
}
func getElementTrail(i int) string {
	return recurs(elements[i], nil) + " trail"
}
func getFabricScrap(i int) string {
	return "Scrap of " + recurs(fabrics[i], nil)
}
func getFoodScraps(i int) string {
	return "Scraps of " + recurs(food[i], nil)
}
func getTextureTrail(i int) string {
	return recurs(textures[i], nil) + " trail"
}
func getWeakness(i int) string {
	return weaknesses[i]
}
func getTactic(i int) string {
	return tactics[i]
}
func getEffectAura(i int) string {
	return recurs(effects[i], nil) + " aura"
}
func getEffectBlast(i int) string {
	return recurs(effects[i], nil) + " blast"
}
func getEffectBolt(i int) string {
	return recurs(effects[i], nil) + " bolt"
}
func getEffectRay(i int) string {
	return recurs(effects[i], nil) + " ray"
}
func getElementControl(i int) string {
	return recurs(elements[i], nil) + " control"
}
func getElementRay(i int) string {
	return recurs(elements[i], nil) + " ray"
}
func getElementBolt(i int) string {
	return recurs(elements[i], nil) + " bolt"
}
func getElementWall(i int) string {
	return recurs(elements[i], nil) + " wall"
}
func getElementBlast(i int) string {
	return recurs(elements[i], nil) + " blast"
}
func getElementAura(i int) string {
	return recurs(elements[i], nil) + " aura"
}
func getPower(i int) string {
	return powers[i]
}
func getElementSkin(i int) string {
	return recurs(elements[i], nil) + " skin"
}
func getMonsterTrait(i int) string {
	return monsterTraits[i]
}
func getAnimalEars(i int) string {
	return recurs(animals[i], nil) + " ears"
}
func getAnimalEyes(i int) string {
	return recurs(animals[i], nil) + " eyes"
}
func getAnimalHead(i int) string {
	return recurs(animals[i], nil) + " head"
}
func getAnimalLimb(i int) string {
	return recurs(animals[i], nil) + " limb"
}
func getAnimalTail(i int) string {
	return recurs(animals[i], nil) + " tail"
}
func getAnimalTeeth(i int) string {
	return recurs(animals[i], nil) + " teeth"
}
func getAnimalVoice(i int) string {
	return recurs(animals[i], nil) + " voice"
}
func getAnimalScent(i int) string {
	return recurs(scents[i], nil) + " scent"
}
func getAnimalBody(i int) string {
	return recurs(animals[i], nil) + " body"
}
func getFoodAddict(i int) string {
	return "Addicted to eating " + recurs(food[i], nil)
}
func getLiability(i int) string {
	return liabilities[i]
}
func getBookExpert(i int) string {
	return "Expert in " + recurs(books[i], nil)
}
func getAsset(i int) string {
	return assets[i]
}
func getProfession(i int) string {
	return professions[i]
}
func getGoal(i int) string {
	return goals[i]
}
func getPersonality(i int) string {
	return personalities[i]
}
func getNPCDetail(i int) string {
	return npcDetails[i]
}
func getName(i int) string {
	if rand.Intn(2) == 0 {
		return maleNames[i]
	}
	return femaleNames[i]
}
func getFactionAlly(i int) string {
	return "Ally " + recurs(factions[i], nil)
}
func getMonsterAlly(i int) string {
	return "Ally " + recurs(monsters[i], nil)
}
func getReward(i int) string {
	return rewards[i]
}
func getMission(i int) string {
	return missions[i]
}
func getFactionWar(i int) string {
	return "War between " + recurs(factions[i], nil) + " and " + recurs(factions[rand.Intn(100)], nil)
}
func getFactionTrait(i int) string {
	return factionsTraits[i]
}
func getFaction(i int) string {
	return factions[i]
}
func getCityEvent(i int) string {
	return cityEvents[i]
}
func getColorSkin(i int) string {
	return recurs(colors[i], nil) + " skin"
}
func getSkinTexture(i int) string {
	return recurs(textures[i], nil) + " skin"
}
func getPotionRecipe(i int) string {
	return "Recipe for a potion of " + recurs(potions[i], nil)
}
func getDisaster(i int) string {
	return disasters[i]
}
func getElementBreath(i int) string {
	return recurs(elements[i], nil) + " breath"
}
func getElementBody(i int) string {
	return recurs(elements[i], nil) + " body"
}
func getElementBlood(i int) string {
	return recurs(elements[i], nil) + " blood"
}
func getDomainFear(i int) string {
	return "Fears " + recurs(domains[i], nil)
}
func getSymbolFear(i int) string {
	return "Fears " + recurs(symbols[i], nil)
}
func getMutation(i int) string {
	return mutations[i]
}
func getDelusion(i int) string {
	return delusions[i]
}
func getHazard(i int) string {
	return hazards[i]
}
func getRandSpell(i int) string {
	return getRandomSpellProxy(false)
}
func getElementField(i int) string {
	return recurs(elements[i], nil) + " field"
}
func getLocation(i int) string {
	return locations[i]
}
func getColorSky(i int) string {
	return recurs(colors[i], nil) + " sky"
}
func getEffectRain(i int) string {
	return recurs(effects[i], nil) + " rain"
}
func getElementRain(i int) string {
	return recurs(elements[i], nil) + " rain"
}
func getFoodRain(i int) string {
	return recurs(food[i], nil) + " rain"
}
func getTasteRain(i int) string {
	return recurs(tastes[i], nil) + " rain"
}
func getTextureBody(i int) string {
	return recurs(textures[i], nil) + " body"
}
func getTextureRain(i int) string {
	return recurs(textures[i], nil) + " rain"
}
func getWeaponRain(i int) string {
	return recurs(weapons[i], nil) + " rain"
}
func getTravelShift(i int) string {
	return travelShifts[i]
}
func getAnimalForm(i int) string {
	return recurs(animals[i], nil) + " form"
}
func getElementForm(i int) string {
	return recurs(elements[i], nil) + " form"
}
func getElementFlow(i int) string {
	return recurs(elements[i], nil) + " flow"
}
func getMechanism(i int) string {
	return mechanisms[i]
}
func getMonsterSkin(i int) string {
	return recurs(monsters[i], nil) + " skin"
}
func getAnimalSkin(i int) string {
	return recurs(animals[i], nil) + " skin"
}
func getRoom(i int) string {
	return rooms[i]
}
func getInn(i int) string {
	innNamePart1 := recurs(innName1[i], nil)
	innNamePart2 := recurs(innName2[rand.Intn(len(innName2))], nil)
	return innNamePart1 + " " + innNamePart2
}
func getDungeon(i int) string {
	return dungeons[i]
}
func getBuilding(i int) string {
	return buildings[i]
}
func getTreasure(i int) string {
	return treasures[i]
}
func getSpellbook(i int) string {
	// return spellbooks[i]
	return "A spellbook, todo"
}
func getSpellScroll(i int) string {
	return "Scroll of " + setSpells[i]
}
func getPotion(i int) string {
	return potions[i]
}
func getIngredient(i int) string {
	return ingredients[i]
}
func getGiantAnimal(i int) string {
	return "Giant " + animals[i]
}
func getDelveShift(i int) string {
	return delveShifts[i]
}
func getMagicSchool(i int) string {
	return magicSchools[i]
}
func getItemTrait(i int) string {
	return itemTraits[i]
}
func getCityTheme(i int) string {
	return cityThemes[i]
}
func getBook(i int) string {
	return books[i]
}
func getArchetype(i int) string {
	return archetypes[i]
}

func getMannerism(i int) string {
	return mannerisms[i]
}
func getRoomTheme(i int) string {
	return roomThemes[i]
}
func getTrapEffect(i int) string {
	return trapEffects[i]
}
func getMonster(i int) string {
	return monsters[i]
}
func getClothing(i int) string {
	return clothing[i]
}
func getMiscItem(i int) string {
	return miscellenousItems[i]
}
func getWeapon(i int) string {
	return weapons[i]
}
func getTool(i int) string {
	return tools[i]
}
func getSymbol(i int) string {
	return symbols[i]
}
func getStructure(i int) string {
	return structures[i]
}
func getOrgan(i int) string {
	return organs[i]
}
func getAnimal(i int) string {
	return animals[i]
}
func getFabric(i int) string {
	return fabrics[i]
}
func getDomain(i int) string {
	return domains[i]
}
func getFood(i int) string {
	return food[i]
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
	return recurs(wizardNames[i], nil)
}
func getSurname(i int) string {
	return surname1[i] + surname2[rand.Intn(len(surname2))]
}
func getTexture(i int) string {
	return textures[i]
}
func getTaste(i int) string {
	return tastes[i]
}
func getSound(i int) string {
	return sounds[i]
}
func getScent(i int) string {
	return scents[i]
}
func getPlaceTrait(i int) string {
	return placeTraits[i]
}
func getColor(i int) string {
	return colors[i]
}
func getMaterials(i int) string {
	return materials[i]
}
