package pirateborg

import (
	"fmt"
	"strings"

	"github.com/genrpg/utils"
)

type Rapscallion struct {
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

func NewRapscallion() PlayerClass {
	rapscallion := &Rapscallion{
		Lvl:         1,
		DevilsLuck:  "d2",
		WeaponDie:   6,
		ClothingDie: 6,
		HatDie:      10,
		StatMods: Stats{
			Strength:  -1,
			Agility:   2,
			Toughness: -1,
		},
		HitDie: 8,
	}
	rapscallion.getStartingFeature()
	return rapscallion
}

func (r *Rapscallion) GetDevilsLuck() string {
	return r.DevilsLuck
}

func (r *Rapscallion) GetClothingDie() int {
	return r.ClothingDie
}
func (r *Rapscallion) GetHatDie() int {
	return r.HatDie
}
func (r *Rapscallion) GetWeaponDie() int {
	return r.WeaponDie
}

func (r *Rapscallion) GetFeatures() []Feature {
	return r.Features
}

func (r *Rapscallion) GetHPDie() int {
	return r.HitDie
}

func (r *Rapscallion) GetStatMods() Stats {
	return r.StatMods
}

func (*Rapscallion) String() string {
	return "Rapscallion"
}

func (r *Rapscallion) StartingFeatureBlurb() string {
	return r.StartingFeature.String()
}

func (r *Rapscallion) Level() int {
	return r.Lvl
}

func (r *Rapscallion) LevelUp() {}

func (r *Rapscallion) GetItems() []Item {
	return r.Items
}

func (r *Rapscallion) GetWeapons() []Weapon {
	return r.Weapons
}

func (r *Rapscallion) Description() string {
	return "A sneaky, cutthroat scallywag good at backstabbing, breaking & entering, stealing, cheating, and escaping. Found in taverns, shadows, and shallow graves." + Red(I("You will need a deck of playing cards..."))
}

type RapsFeature struct {
	Name        string
	Description string
}

func (rf RapsFeature) String() string {
	return fmt.Sprintf("%s | %s", strings.ToUpper(rf.Name), rf.Description)
}

func (r *Rapscallion) getStartingFeature() {
	rf := rapscallionStartingFeatures[utils.D(6)]

	r.StartingFeature = rf
	// r.Features = append(r.Features, RapsFeature{Name: "Drinking Grog & Rum",
	// 	Description: "Test TOUGHNESS DR8 + [number of drinks in the last hour] to heal " + B("d4") + " HP. Fail and you vomit for " + B("d2") + " rounds. Agility is " + B("-1") + " for each drink (lasts one hour per drink)."})
	r.Features = append(r.Features, rf)
}

type StartingRapscallionFeature RapsFeature

func (rf StartingRapscallionFeature) String() string {
	return fmt.Sprintf("%s | %s", strings.ToUpper(rf.Name), rf.Description)
}

var rapscallionStartingFeatures = map[int]StartingRapscallionFeature{1: rapscallionFeature1, 2: rapscallionFeature2, 3: rapscallionFeature3, 4: rapscallionFeature4, 5: rapscallionFeature5, 6: rapscallionFeature6}

var rapscallionFeature1 = StartingRapscallionFeature{
	Name:        "Back Stabber",
	Description: "If you attack by surprise " + I("(from hiding, distracted enemy, etc.)") + " lower the attack DR by 2 and deal " + B("d2") + " extra damage.",
}

var rapscallionFeature2 = StartingRapscallionFeature{
	Name:        "Burglar",
	Description: "You begin with lock picks. Pickpocket, disarm an enemy, or disable a trap: -4 to DR.",
}

var rapscallionFeature3 = StartingRapscallionFeature{
	Name:        "Rope Monkey",
	Description: "You're as nimble in the rigging as you are on the deck. If you attack after swinging, jumping, or making an acrobatic maneuver, test AGILITY DR10 to automatically hit and deal +2 damage.",
}

var rapscallionFeature4 = StartingRapscallionFeature{
	Name:        "Sneaky Bastard",
	Description: "When striking from the shadows or while sneaking, test AGILITY DR12. On a success you automatically deal a critical hit.",
}

var rapscallionFeature5 = StartingRapscallionFeature{
	Name:        "Lucky Devil",
	Description: "Whenever you use the " + I(B("Devil's Luck")) + ", draw a card: 9+: Regain 1 Luck. Joker: Roll on the " + B("Joker Table") + I(" (page 37)") + ".",
}

var rapscallionFeature6 = StartingRapscallionFeature{
	Name:        "Grog Brewer",
	Description: "Each day you can brew " + B("d4") + " servings of potent grog. You can soak melee weapons in it to use as a poison Grog in a wound: Test TOUGHNESS DR14 or " + B("-d6") + " HP",
}
