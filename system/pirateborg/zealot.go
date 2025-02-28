package pirateborg

import (
	"fmt"

	"github.com/genrpg/utils"
	"github.com/ttacon/chalk"
)

type Zealot struct {
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

func NewZealot() PlayerClass {
	zealot := &Zealot{
		Lvl:         1,
		DevilsLuck:  "d4",
		WeaponDie:   8,
		ClothingDie: 8,
		HatDie:      0,
		StatMods: Stats{
			Spirit:    2,
			Agility:   -1,
			Toughness: -1,
		},
		HitDie: 8,
	}
	zealot.getStartingFeature()
	return zealot
}

func (z *Zealot) GetDevilsLuck() string {
	return z.DevilsLuck
}

func (z *Zealot) GetClothingDie() int {
	return z.ClothingDie
}
func (z *Zealot) GetHatDie() int {
	return z.HatDie
}
func (z *Zealot) GetWeaponDie() int {
	return z.WeaponDie
}

func (z *Zealot) GetFeatures() []Feature {
	return z.Features
}

func (z *Zealot) GetHPDie() int {
	return z.HitDie
}

func (z *Zealot) GetStatMods() Stats {
	return z.StatMods
}

func (*Zealot) String() string {
	return "Zealot"
}

func (z *Zealot) StartingFeatureBlurb() string {
	return z.StartingFeature.String()
}

func (z *Zealot) Level() int {
	return z.Lvl
}

func (z *Zealot) LevelUp() {}

func (z *Zealot) GetItems() []Item {
	return z.Items
}

func (z *Zealot) GetWeapons() []Weapon {
	return z.Weapons
}

func (z *Zealot) Description() string {
	return "A clergy member, cultist, shaman, or believer."
}

type ZealotFeature struct {
	Name        string
	Description string
}

func (zf ZealotFeature) String() string {
	return fmt.Sprintf("%s | %s", zf.Name, zf.Description)
}

func (z *Zealot) getStartingFeature() {
	zf := zealotStartingFeatures[utils.D(6)]

	z.StartingFeature = zf
	z.Features = append(z.Features, ZealotFeature{Name: "ARMORED CASTER",
		Description: "Thou may use ancient relics and arcane rituals whilst wearing medium armor " + I("(tier 2 or lower)") + "."})
	z.Features = append(z.Features, ZealotFeature{Name: "Whom dost thou serve?", Description: []string{
		"THE ONE TRUE GOD", "MOTHER NATURE", "THE ANCIENT GODS", "CHAOS", "THE DEEP", "THE DARK ONE", "THE CHURCH", "THE GREAT OLD ONE",
	}[utils.TableDie(8)]})
	z.Features = append(z.Features, ZealotFeature{Name: "PRAYING " + I(chalk.Magenta.Color("SPELLCASTING")),
		Description: "Thou may use prayers " + B("d2+SPIRIT") + " times a day (reroll when dawn breaks). It consumes thy action to pray, but doth not require a roll or test.\nSpells:"})

	z.Features = append(z.Features, zf)
}

type StartingZealotFeature ZealotFeature

func (zf StartingZealotFeature) String() string {
	return fmt.Sprintf("%s | %s", zf.Name, zf.Description)
}

var zealotStartingFeatures = map[int]StartingZealotFeature{1: zealotFeature1, 2: zealotFeature2, 3: zealotFeature3, 4: zealotFeature4, 5: zealotFeature5, 6: zealotFeature6,
	7: zealotFeature7, 8: zealotFeature8, 9: zealotFeature9, 10: zealotFeature10}

var zealotFeature1 = StartingZealotFeature{
	Name:        chalk.Magenta.Color("HEAL"),
	Description: "Heal thyself or another for " + B("d8") + " HP.",
}

var zealotFeature2 = StartingZealotFeature{
	Name:        chalk.Magenta.Color("CURSE"),
	Description: "Test SPIRIT DR10: deal " + B("d8+SPIRIT") + " damage to an enemy that thou cannot see. DR8 if it has already been hurt in this fight.",
}

var zealotFeature3 = StartingZealotFeature{
	Name:        chalk.Magenta.Color("DEATH WARD"),
	Description: "Touch the corpse of one who hath just died and test SPIRIT DR10: they return to life with 1 HP. Crit: Full HP. Fumble: They come back as a zombie and attacketh thee!",
}

var zealotFeature4 = StartingZealotFeature{
	Name:        chalk.Magenta.Color("CONTROL WEATHER"),
	Description: "Test SPIRIT DR10 to change the direction of the wind. If thou succeedeth by 5 or more, thou can also conjure or repel precipitation. Crit: Lightning striketh thine enemy, " + B("d12") + ". Fumble: Lightning strikes thee for " + B("d6") + ".",
}

var zealotFeature5 = StartingZealotFeature{
	Name:        chalk.Magenta.Color("BLESSED GUIDANCE"),
	Description: "Thou may add " + B("d4") + " to any roll thee or another player maketh. Use this at any time, including after a roll (does not taketh thy action).",
}

var zealotFeature6 = StartingZealotFeature{
	Name:        chalk.Magenta.Color("HOLY PROTECTION"),
	Description: "Thou or thine ally gets -4 to DRs to defend for one attack. Use this at any time, including after a roll (does not taketh thy action).",
}

var zealotFeature7 = StartingZealotFeature{
	Name:        chalk.Magenta.Color("DIVINE LIGHT"),
	Description: "Bright light radiates from thee for up to " + B("d6") + " x 10 minutes. Enemies that see it are -2 DR to defend against.",
}

var zealotFeature8 = StartingZealotFeature{
	Name:        chalk.Magenta.Color("SILENCE"),
	Description: "For the next " + B("2d6") + " x 10 minutes, everything within 25' of thee maketh no sound. The effect only ends when the time doth expire.",
}

var zealotFeature9 = StartingZealotFeature{
	Name:        chalk.Magenta.Color("SANCTUARY"),
	Description: "All thy brethren in sight heal " + B("d4") + " HP.",
}

var zealotFeature10 = StartingZealotFeature{
	Name:        chalk.Magenta.Color("COMMUNE"),
	Description: "Test SPIRIT DR8: Asketh thy deity a singular \"yay\" or \"nay\" query. Thy response may be \"unclear\" or thou may receiveth no answer.",
}
