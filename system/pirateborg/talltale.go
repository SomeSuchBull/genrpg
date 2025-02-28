package pirateborg

import (
	"fmt"

	"github.com/genrpg/utils"
)

type TallTale struct {
	StatMods        Stats
	Features        []Feature `json:"features"`
	AdditionalClass PlayerClass
	Descrptn        string `json:"description"`
	Name            string `json:"name"`
	DevilsLuck      string
	Die             int
	HitDie          int
	Lvl             int
	WeaponDie       int
}

func NewTallTale() PlayerClass {
	tallTale := &TallTale{
		AdditionalClass: GetClass(),
		Lvl:             1,
	}
	tallTale.Die = utils.D(3)
	switch tallTale.Die {
	case 1:
		tallTale.GetMerfolk()
	case 2:
		tallTale.GetAquaticMutant()
	case 3:
		tallTale.GetSentientAnimal()
	}
	return tallTale
}

func (tt *TallTale) GetDevilsLuck() string {
	if tt.Die == 3 {
		return tt.DevilsLuck
	}
	return tt.AdditionalClass.GetDevilsLuck()
}

func (tt *TallTale) GetClothingDie() int {
	if tt.Die == 3 {
		return 0
	}
	return tt.AdditionalClass.GetClothingDie()
}
func (tt *TallTale) GetHatDie() int {
	if tt.Die == 3 {
		return 0
	}
	return tt.AdditionalClass.GetHatDie()
}
func (tt *TallTale) GetWeaponDie() int {
	if tt.Die == 3 {
		return tt.WeaponDie
	}
	return tt.AdditionalClass.GetWeaponDie()
}

func (tt *TallTale) GetFeatures() []Feature {
	if tt.Die == 3 {
		return tt.Features
	}
	return append(tt.Features, tt.AdditionalClass.GetFeatures()...)
}

func (tt *TallTale) GetHPDie() int {
	if tt.Die == 3 {
		return tt.HitDie
	}
	return tt.AdditionalClass.GetHPDie()
}

func (tt *TallTale) GetStatMods() Stats {
	if tt.Die == 3 {
		return tt.StatMods
	}
	return tt.AdditionalClass.GetStatMods()
}

func (tt *TallTale) String() string {
	if tt.Die == 3 {
		return tt.Name
	}
	return fmt.Sprintf("%s | %s", tt.Name, tt.AdditionalClass.String())
}

func (tt *TallTale) StartingFeatureBlurb() string {
	return ""
}

func (tt *TallTale) Level() int {
	if tt.Die == 3 {
		return tt.Lvl
	}
	return tt.AdditionalClass.Level()
}

func (tt *TallTale) LevelUp() {
	if tt.Die == 3 {
		return
	}
	tt.AdditionalClass.LevelUp()
}

func (tt *TallTale) GetItems() []Item {
	if tt.Die == 3 {
		return []Item{}
	}
	return tt.AdditionalClass.GetItems()
}

func (tt *TallTale) GetWeapons() []Weapon {
	if tt.Die == 3 {
		return []Weapon{}
	}
	return tt.AdditionalClass.GetWeapons()
}

func (tt *TallTale) Description() string {
	if tt.Die == 3 {
		return tt.Descrptn
	}
	return fmt.Sprintf("%s\n%s", tt.Descrptn, tt.AdditionalClass.Description())
}

type TTFeature struct {
	Name        string
	Description string
}

func (ttf TTFeature) String() string {
	return fmt.Sprintf("%s | %s", ttf.Name, ttf.Description)
}

func (tt *TallTale) GetMerfolk() {
	tt.Name = "Merfolk"
	tt.Descrptn = "One of the children of Poseidon, your lower half is fish-like and you have gills. Both can be magically hidden as you masquerade as a human."
	tt.Features = []Feature{
		TTFeature{Name: "Water Habitation", Description: "Lower all DRs by 4 when underwater. You die if you go 1+TOUGHNESS days without submerging in fresh seawater (minimum 1 day)"},
	}
}

func (tt *TallTale) GetAquaticMutant() {
	tt.Name = "Aquatic Mutant"
	tt.Descrptn = "You are the spawn from an abominable union between man and the sea."
	var ttf TTFeature
	switch utils.D(8) {
	case 1:
		ttf = TTFeature{Name: "Anglerfish", Description: "You see in the dark. Bite: " + B("d4")}
	case 2:
		ttf = TTFeature{Name: "Crab", Description: "Pincher: " + B("d6")}
	case 3:
		ttf = TTFeature{Name: "Jellyfish", Description: "Tendril: " + B("d2") + " & stun 1 round, 10' reach"}
	case 4:
		ttf = TTFeature{Name: "Octopus", Description: "Tentacles: " + B("2d4")}
	case 5:
		ttf = TTFeature{Name: "Sea Turtle", Description: "Extra -d2 armor"}
	case 6:
		ttf = TTFeature{Name: "Electric Eel", Description: "+1 AGILITY. Electric skin: " + B("d6")}
	case 7:
		ttf = TTFeature{Name: "Shark", Description: "Bite: " + B("d8")}
	case 8:
		ttf = TTFeature{Name: "The Great Old One", Description: "Learn one random Ritual (pg. 64)."}
	}
	tt.Features = []Feature{ttf}
}

func (tt *TallTale) GetSentientAnimal() {
	tt.Descrptn = "You are a mystically intelligent animal. Maybe you were human once.\nWhen you begin:\n— Reroll any backstory details that don't make sense.\n— Equipment you can't conceivably carry can be given away or left behind."
	switch utils.D(6) {
	case 1:
		tt.Name = "Foul Fowl"
		tt.Features = []Feature{TTFeature{Name: "Foul Fowl", Description: "When you are killed, the ghosts of a hundred chickens swarm your assailant, ripping their spiritual soul from their flesh."},
			TTFeature{Name: "Magic", Description: "Gain the ability from one random Relic (p. 62) & one random Ritual (p. 64)."},
			TTFeature{Name: "Beak Peck", Description: B("d2") + " damage."}}
		tt.StatMods = Stats{Strength: -2, Agility: -2, Toughness: -2, Presence: -2, Spirit: 3}
		tt.HitDie = 4
		tt.DevilsLuck = "d4"
	case 2:
		tt.Name = "Jaguar"
		tt.Features = []Feature{TTFeature{Name: "Jaguar", Description: "You're a deadly jungle cat ."},
			TTFeature{Name: "Bite/Claws", Description: B("d8") + " damage."}}
		tt.StatMods = Stats{Strength: 2, Agility: 2, Toughness: -2, Presence: -2, Spirit: -2}
		tt.HitDie = 8
		tt.DevilsLuck = "d4"
	case 3:
		tt.Name = "Crocodile"
		tt.Features = []Feature{TTFeature{Name: "Crocodile", Description: "You can swim and hide well in water."},
			TTFeature{Name: "Bite", Description: B("d10") + " damage."}}
		tt.StatMods = Stats{Strength: 3, Agility: -2, Toughness: 1, Presence: -2, Spirit: -2}
		tt.HitDie = 10
		tt.DevilsLuck = "d4"
	case 4:
		tt.Name = "Bilge rat"
		tt.Features = []Feature{TTFeature{Name: "Bilge rat", Description: "You are a filthy rodent."},
			TTFeature{Name: "Diseased Bite", Description: B("d2") + " damage. 1-in-6 chance the target dies in " + B("d4") + " rounds from whatever disease you are carrying."}}
		tt.StatMods = Stats{Strength: -2, Agility: 3, Toughness: 2, Presence: -2, Spirit: -2}
		tt.HitDie = 2
		tt.DevilsLuck = "d4"
	case 5:
		tt.Name = "Lucky Parrot"
		tt.Features = []Feature{TTFeature{Name: "Lucky Parrot", Description: "A colorful, talking, flying bird"},
			TTFeature{Name: "Beak Peck", Description: B("d4") + " damage."}}
		tt.StatMods = Stats{Strength: -2, Agility: 1, Toughness: -2, Presence: 2, Spirit: -2}
		tt.HitDie = 2
		tt.DevilsLuck = "d6"
	case 6:
		tt.Name = "Clever Monkey"
		tt.Features = []Feature{TTFeature{Name: "Clever Monkey", Description: "You're an excellent climber."}}
		tt.StatMods = Stats{Strength: -1, Agility: +2, Toughness: -2, Spirit: -2}
		tt.HitDie = 6
		tt.DevilsLuck = "d4"
		tt.WeaponDie = 10
	}
}
