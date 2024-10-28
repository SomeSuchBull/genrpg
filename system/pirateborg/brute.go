package pirateborg

import (
	"fmt"
	"strings"

	"github.com/genrpg/utils"
)

type Brute struct {
	Lvl             int       `json:"level"`
	HitDie          int       `json:"hitDie"`
	StatMods        Stats     `json:"statMods"`
	DevilsLuck      string    `json:"devilsLuck"`
	ClothingDie     int       `json:"clothingDie"`
	HatDie          int       `json:"hatDie"`
	StartingFeature Feature   `json:"startingFeature"`
	Features        []Feature `json:"features"`
	Items           []Item    `json:"item"`
	Weapons         []Weapon  `json:"weapons"`
}

func NewBrute() PlayerClass {
	brute := &Brute{
		Lvl:         1,
		DevilsLuck:  "d2",
		ClothingDie: 10,
		HatDie:      12,
		StatMods: Stats{
			Strength:  1,
			Toughness: 1,
			Presence:  -1,
			Spirit:    -1,
		},
		HitDie: 12,
	}
	brute.getStartingFeature()
	return brute
}

func (b *Brute) GetDevilsLuck() string {
	return b.DevilsLuck
}

func (b *Brute) GetClothingDie() int {
	return b.ClothingDie
}
func (b *Brute) GetHatDie() int {
	return b.HatDie
}
func (b *Brute) GetWeaponDie() int {
	return 0
}

func (b *Brute) GetFeatures() []Feature {
	return b.Features
}

func (b *Brute) GetHPDie() int {
	return b.HitDie
}

func (b *Brute) GetStatMods() Stats {
	return b.StatMods
}

func (*Brute) String() string {
	return "Brute"
}

func (b *Brute) StartingFeatureBlurb() string {
	return b.StartingFeature.String()
}

func (b *Brute) Level() int {
	return b.Lvl
}

// TODO: decide if this is something even worth pursuing
func (b *Brute) LevelUp() {}

func (b *Brute) GetItems() []Item {
	return b.Items
}

func (b *Brute) GetWeapons() []Weapon {
	return b.Weapons
}

func (b *Brute) Description() string {
	return "When you're not bashing, smashing, slashing, or crashing, you're... well, that's all you really know how to do. You can’t use " + B("Arcane Rituals") + ", but your muscles are basically \"magic.\""
}

func (b *Brute) getStartingFeature() {
	sf := bruteStartingFeatures[utils.D(6)]
	b.StartingFeature = sf
	b.Weapons = append(b.Weapons, Weapon(sf))
	// b.Features = append(b.Features, RapsFeature{Name: "Drinking Grog & Rum",
	// 	Description: "Test TOUGHNESS DR8 + [number of drinks in the last hour] to heal " + B("d4") + " HP. Fail and you vomit for " + B("d2") + " rounds. Agility is " + B("-1") + " for each drink (lasts one hour per drink)."})
	b.Features = append(b.Features, sf)
}

type StartingBruteFeature Weapon

func (bf StartingBruteFeature) String() string {
	return fmt.Sprintf("%s | %s | %s", strings.ToUpper(bf.Name), bf.Damage, bf.Extra)
}

var bruteStartingFeatures = map[int]StartingBruteFeature{1: bruteFeature1, 2: bruteFeature2, 3: bruteFeature3, 4: bruteFeature4, 5: bruteFeature5, 6: bruteFeature6}

var bruteFeature1 = StartingBruteFeature{
	Item: Item{
		Name:  "Brass anchor",
		Extra: "Requires 2 hands and target's armor is reduced by one tier " + B("(-d2)") + " during the attack.",
	},
	Damage: "d8",
}

var bruteFeature2 = StartingBruteFeature{
	Item: Item{
		Name:  "Whaling harpoon",
		Extra: "Can be thrown by testing " + B("AGILITY DR10") + ".",
	},
	Damage: "d8",
}

var bruteFeature3 = StartingBruteFeature{
	Item: Item{
		Name:  "Meat cleaver",
		Extra: "Caked in layers of dried blood. On a damage roll of " + B("1") + " it spreads a disease from one of its prior victims: the target loses " + B("d6") + " HP at the start of its next " + B("two") + " turns.",
	},
	Damage: "d4",
}

var bruteFeature4 = StartingBruteFeature{
	Item: Item{
		Name:  "Part of a broken mast",
		Extra: "It has a rusted nail protruding from one end. Deals an extra " + B("d6") + " on a critical hit.",
	},
	Damage: "d8",
}

var bruteFeature5 = StartingBruteFeature{
	Item: Item{
		Name:  "Runic machete",
		Extra: "Great for chopping down vines and fopdoodles. It glows in the dark if there are undead nearby.",
	},
	Damage: "d6",
}

var bruteFeature6 = StartingBruteFeature{
	Item: Item{
		Name:  "Rotten cargo net",
		Extra: "Test " + B("AGILITY DR12") + " to throw it at something and stop it from moving for " + B("d2") + " rounds. Trapped targets take " + B("d2") + " damage/round.",
	},
	Damage: "-",
}
