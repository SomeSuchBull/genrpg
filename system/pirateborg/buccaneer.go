package pirateborg

import (
	"fmt"
	"strings"

	"github.com/genrpg/utils"
)

type Buccaneer struct {
	StatMods        Stats     `json:"statMods"`
	Features        []Feature `json:"features"`
	Items           []Item    `json:"item"`
	Weapons         []Weapon  `json:"weapons"`
	StartingFeature Feature   `json:"startingFeature"`
	DevilsLuck      string    `json:"devilsLuck"`
	Lvl             int       `json:"level"`
	HitDie          int       `json:"hitDie"`
	WeaponDie       int       `json:"weaponDie"`
	ClothingDie     int       `json:"clothingDie"`
	HatDie          int       `json:"hatDie"`
}

func NewBuccaneer() PlayerClass {
	buccaneer := &Buccaneer{
		Lvl:         1,
		DevilsLuck:  "d2",
		WeaponDie:   0,
		ClothingDie: 10,
		HatDie:      12,
		StatMods: Stats{
			Presence: 2,
			Agility:  -1,
			Spirit:   -1,
		},
		HitDie: 8,
	}
	buccaneer.getStartingFeature()
	return buccaneer
}

func (bucc *Buccaneer) GetDevilsLuck() string {
	return bucc.DevilsLuck
}

func (bucc *Buccaneer) GetClothingDie() int {
	return bucc.ClothingDie
}
func (bucc *Buccaneer) GetHatDie() int {
	return bucc.HatDie
}
func (bucc *Buccaneer) GetWeaponDie() int {
	return bucc.WeaponDie
}

func (bucc *Buccaneer) GetFeatures() []Feature {
	return bucc.Features
}

func (bucc *Buccaneer) GetHPDie() int {
	return bucc.HitDie
}

func (bucc *Buccaneer) GetStatMods() Stats {
	return bucc.StatMods
}

func (*Buccaneer) String() string {
	return "Buccaneer"
}

func (bucc *Buccaneer) StartingFeatureBlurb() string {
	return bucc.StartingFeature.String()
}

func (bucc *Buccaneer) Level() int {
	return bucc.Lvl
}

func (bucc *Buccaneer) LevelUp() {}

func (bucc *Buccaneer) GetItems() []Item {
	return bucc.Items
}

func (bucc *Buccaneer) GetWeapons() []Weapon {
	return []Weapon{{Item: Item{Name: "Musket", Extra: "reload 1 action, range 150' " + I("(10 + PRESENCE rounds of shot)")}, Damage: "2d6"}}
}

func (bucc *Buccaneer) Description() string {
	return "Skilled trackers and survivalists. Expert sharpshooters, especially with muskets & rifles."
}

type BuccFeature struct {
	Name        string
	Description string
}

func (buccf BuccFeature) String() string {
	return fmt.Sprintf("%s | %s", strings.ToUpper(buccf.Name), buccf.Description)
}

func (bucc *Buccaneer) getStartingFeature() {
	buccf := buccaneerStartingFeatures[utils.D(6)]

	bucc.StartingFeature = buccf
	// bucc.Features = append(bucc.Features, BuccFeature{Name: "Drinking Grog & Rum",
	// 	Description: "Test TOUGHNESS DR8 + [number of drinks in the last hour] to heal " + B("d4") + " HP. Fail and you vomit for " + B("d2") + " rounds. Agility is " + B("-1") + " for each drink (lasts one hour per drink)."})
	bucc.Features = append(bucc.Features, BuccFeature{Name: "Marksman",
		Description: "Reloading " + B("black powder weapons") + " takes you " + B("1") + " round instead of " + B("2") + "."})
	bucc.Features = append(bucc.Features, buccf)
}

type StartingBuccaneerFeature BuccFeature

func (buccf StartingBuccaneerFeature) String() string {
	return fmt.Sprintf("%s | %s", strings.ToUpper(buccf.Name), buccf.Description)
}

var buccaneerStartingFeatures = map[int]StartingBuccaneerFeature{1: buccaneerFeature1, 2: buccaneerFeature2, 3: buccaneerFeature3, 4: buccaneerFeature4, 5: buccaneerFeature5, 6: buccaneerFeature6}

var buccaneerFeature1 = StartingBuccaneerFeature{
	Name:        "Treasure Hunter",
	Description: "Ability tests related to mapping, navigating, treasure hunting, finding & disarming traps, and tracking prey are -3 DR.",
}

var buccaneerFeature2 = StartingBuccaneerFeature{
	Name:        "Crack Shot",
	Description: "All ranged attacks are -2 DR.",
}

var buccaneerFeature3 = StartingBuccaneerFeature{
	Name:        "Fix Bayonets!",
	Description: "You now have a bayonet (" + B("d4") + "). You can attack with it on the same turn you reload.",
}

var buccaneerFeature4 = StartingBuccaneerFeature{
	Name:        "Focused Aim",
	Description: "Attacks against enemies you have already shot at during this combat are -4 DR to hit.",
}

var buccaneerFeature5 = StartingBuccaneerFeature{
	Name:        "Buccan Cook",
	Description: "Months of island life have taught you the art of cooking meats over the buccan fire. You start with " + B("d8") + " rations of exquisite smoked meat. Eating it immediately recovers " + B("d4") + " HP, and you can make d4 more rations from any edible animal you kill.",
}

var buccaneerFeature6 = StartingBuccaneerFeature{
	Name:        "Survivalist",
	Description: "Your body has developed into a finely tuned machine for existing in the wild. Gain +1 TOUGHNESS. You cannot become infected, sick, or poisoned, and your maximum HP increases by " + B("d4") + ".",
}
