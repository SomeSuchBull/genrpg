package pirateborg

import (
	"fmt"
	"strings"

	"github.com/genrpg/utils"
)

type Swashbuckler struct {
	Lvl             int       `json:"level"`
	HitDie          int       `json:"hitDie"`
	StatMods        Stats     `json:"statMods"`
	DevilsLuck      string    `json:"devilsLuck"`
	WeaponDie       int       `json:"weaponDie"`
	ClothingDie     int       `json:"clothingDie"`
	HatDie          int       `json:"hatDie"`
	StartingFeature Feature   `json:"startingFeature"`
	Features        []Feature `json:"features"`
	Items           []Item    `json:"item"`
	Weapons         []Weapon  `json:"weapons"`
}

func (pc *PlayerCharacter) NewSwashbuckler() {
	swashbuckler := &Swashbuckler{
		Lvl:         1,
		DevilsLuck:  "d2",
		WeaponDie:   10,
		ClothingDie: 10,
		HatDie:      12,
		StatMods: Stats{
			Strength: 1,
			Agility:  1,
			Presence: -1,
			Spirit:   -1,
		},
		HitDie: 10,
	}
	swashbuckler.getStartingFeature()
	pc.Class = swashbuckler
}

func (sb *Swashbuckler) GetDevilsLuck() string {
	return sb.DevilsLuck
}

func (sb *Swashbuckler) GetClothingDie() int {
	return sb.ClothingDie
}
func (sb *Swashbuckler) GetHatDie() int {
	return sb.HatDie
}
func (sb *Swashbuckler) GetWeaponDie() int {
	return sb.WeaponDie
}

func (sb *Swashbuckler) GetFeatures() []Feature {
	return sb.Features
}

func (sb *Swashbuckler) GetHPDie() int {
	return sb.HitDie
}

func (sb *Swashbuckler) GetStatMods() Stats {
	return sb.StatMods
}

func (*Swashbuckler) String() string {
	return "Swashbuckler"
}

func (sb *Swashbuckler) StartingFeatureBlurb() string {
	return sb.StartingFeature.String()
}

func (sb *Swashbuckler) Level() int {
	return sb.Lvl
}

// TODO: decide if this is something even worth pursuing
func (sb *Swashbuckler) LevelUp() {}

func (sb *Swashbuckler) GetItems() []Item {
	return sb.Items
}

func (sb *Swashbuckler) GetWeapons() []Weapon {
	return sb.Weapons
}

func (sb *Swashbuckler) Description() string {
	return "A brash fighter with " + B(I("Bravado")) + " & " + B(I("Swagger")) + "."
}

type SwashFeature struct {
	Name        string
	Description string
}

func (rf SwashFeature) String() string {
	return fmt.Sprintf("%s | %s", strings.ToUpper(rf.Name), rf.Description)
}

func (sb *Swashbuckler) getStartingFeature() {
	rf := swashbucklerStartingFeatures[utils.D(6)]

	sb.StartingFeature = rf
	// sb.Features = append(sb.Features, SwashFeature{Name: "Drinking Grog & Rum",
	// 	Description: "Test TOUGHNESS DR8 + [number of drinks in the last hour] to heal " + B("d4") + " HP. Fail and you vomit for " + B("d2") + " rounds. Agility is " + B("-1") + " for each drink (lasts one hour per drink)."})
	sb.Features = append(sb.Features, rf)
}

type StartingSwashbucklerFeature SwashFeature

func (rf StartingSwashbucklerFeature) String() string {
	return fmt.Sprintf("%s | %s", strings.ToUpper(rf.Name), rf.Description)
}

var swashbucklerStartingFeatures = map[int]StartingSwashbucklerFeature{1: swashbucklerFeature1, 2: swashbucklerFeature2, 3: swashbucklerFeature3, 4: swashbucklerFeature4, 5: swashbucklerFeature5, 6: swashbucklerFeature6}

var swashbucklerFeature1 = StartingSwashbucklerFeature{
	Name:        "Ostentatious Fencer",
	Description: "Your melee Attack/Defense is DR10 when wielding a " + B("rapier") + " or " + B("cutlass") + ". When dueling one-on-one, you deal +1 damage.",
}

var swashbucklerFeature2 = StartingSwashbucklerFeature{
	Name:        "Flintlock Fanatic",
	Description: "You can attack with up to three pistols on your turn (if you have them). Reloading one pistol only takes you one round.",
}

var swashbucklerFeature3 = StartingSwashbucklerFeature{
	Name:        "Scurvy Scallywag",
	Description: "You don't fight fair. -2 DR when attacking an enemy that has already been attacked this turn.",
}

var swashbucklerFeature4 = StartingSwashbucklerFeature{
	Name:        "Inspiring Leader",
	Description: "Once each combat, roll a " + B("d4") + ". Each of your allies may add or subtract that value from any one roll during this combat.",
}

var swashbucklerFeature5 = StartingSwashbucklerFeature{
	Name:        "Knife Knave",
	Description: "You start with 2 knives (" + B("d4") + ", " + I("add them to your equipment") + "), and when attacking with them you can make two attacks a turn. They are DR10 to hit, and if the first attack hits, the 2nd is an auto-hit.",
}

var swashbucklerFeature6 = StartingSwashbucklerFeature{
	Name:        "Black Powder Poet",
	Description: "You start with explosives. Roll " + B("d4") + " times on the Bombs table " + I("(pg. 53)") + ". Your DR is -2 when throwing bombs.",
}
