package pirateborg

import (
	"fmt"

	"github.com/genrpg/utils"
	"github.com/ttacon/chalk"
)

type Sorcerer struct {
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

func NewSorcerer() PlayerClass {
	sorcerer := &Sorcerer{
		Lvl:         1,
		DevilsLuck:  "d4",
		WeaponDie:   0,
		ClothingDie: 6,
		HatDie:      0,
		StatMods: Stats{
			Spirit:    2,
			Strength:  -1,
			Toughness: -1,
		},
		HitDie: 8,
	}
	sorcerer.getStartingFeature()
	return sorcerer
}

func (s *Sorcerer) GetDevilsLuck() string {
	return s.DevilsLuck
}

func (s *Sorcerer) GetClothingDie() int {
	return s.ClothingDie
}
func (s *Sorcerer) GetHatDie() int {
	return s.HatDie
}
func (s *Sorcerer) GetWeaponDie() int {
	return s.WeaponDie
}

func (s *Sorcerer) GetFeatures() []Feature {
	return s.Features
}

func (s *Sorcerer) GetHPDie() int {
	return s.HitDie
}

func (s *Sorcerer) GetStatMods() Stats {
	return s.StatMods
}

func (*Sorcerer) String() string {
	return "Sorcerer"
}

func (s *Sorcerer) StartingFeatureBlurb() string {
	return s.StartingFeature.String()
}

func (s *Sorcerer) Level() int {
	return s.Lvl
}

func (s *Sorcerer) LevelUp() {}

func (s *Sorcerer) GetItems() []Item {
	return s.Items
}

func (s *Sorcerer) GetWeapons() []Weapon {
	return []Weapon{{Item: Item{Name: "wooden knife or belaying pin."}, Damage: "d4"}}
}

func (s *Sorcerer) Description() string {
	return "An eldritch occultist, frail and enigmatic. They are conduits for meddlesome natural spirits and devious necromantic entities."
}

type SorcererFeature struct {
	Name        string
	Description string
}

func (sf SorcererFeature) String() string {
	return fmt.Sprintf("%s | %s", sf.Name, sf.Description)
}

func (s *Sorcerer) getStartingFeature() {
	sf := sorcererStartingFeatures[utils.D(6)]

	s.StartingFeature = sf
	s.Features = append(s.Features, SorcererFeature{Name: "CASTER",
		Description: "You can use ancient relics and arcane rituals when wearing medium armor, but NEVER while near cold iron or while holding metal."})
	s.Features = append(s.Features, SorcererFeature{Name: chalk.Magenta.Color("SPELLCASTING"),
		Description: "You can cast " + B("d2+SPIRIT") + " number of spells each day, resetting at sunset. They take your action to cast, but do not require a roll or test.\nSpells:"})

	s.Features = append(s.Features, sf)
}

type StartingSorcererFeature SorcererFeature

func (sf StartingSorcererFeature) String() string {
	return fmt.Sprintf("%s | %s", sf.Name, sf.Description)
}

var sorcererStartingFeatures = map[int]StartingSorcererFeature{1: sorcererFeature1, 2: sorcererFeature2, 3: sorcererFeature3, 4: sorcererFeature4, 5: sorcererFeature5, 6: sorcererFeature6}

var sorcererFeature1 = StartingSorcererFeature{
	Name:        chalk.Magenta.Color("DEAD HEAD"),
	Description: "You summon a flying, ghostly skull. You may spend your action and test SPIRIT DR12 to have it deal damage to 1 target. It dissipates after 1 minute or if it deals any damage. Fumble: it attacks you. It ignores armor and deals " + B("2d4") + ".",
}

var sorcererFeature2 = StartingSorcererFeature{
	Name:        chalk.Magenta.Color("SPIRITUAL POSSESSION"),
	Description: "One random creature is possessed by a spirit or ghost. Ally: -2 DR to attack and defense. Enemy: -2 DR to attack or defend against it. Any Fumbles related to this creature cause the spirit to leave, stunning the host for 1 round. Lasts for " + B("d2") + " rounds",
}

var sorcererFeature3 = StartingSorcererFeature{
	Name:        chalk.Magenta.Color("PROTECTION"),
	Description: "You summon a ghost or spirit to watch over the souls of you and your allies. Everyone who is protected gets -d2 protection for one hour as if wearing extra armor (does not affect penalties to Strength and Agility, not affected by Fumbles). 1 soul",
}

var sorcererFeature4 = StartingSorcererFeature{
	Name:        chalk.Magenta.Color("CLAIRVOYANCE"),
	Description: "Ask the spirits a question about an adjacent room or area, though their answer may be a lie. Test SPIRIT DR12 to know if they are telling the truth.",
}

var sorcererFeature5 = StartingSorcererFeature{
	Name:        chalk.Magenta.Color("NECRO-SLEEP"),
	Description: "A living creature appears to fall over dead, but when they awake they remember everything. Test SPIRIT DR12 to see if it falls \"dead\" asleep for " + B("d2") + " rounds",
}

var sorcererFeature6 = StartingSorcererFeature{
	Name:        chalk.Magenta.Color("RAISE THE DEAD"),
	Description: "You can create skeletal thralls from nearby corpses. They are stupid, but obey your verbal commands. They tumble into bones at sunrise. " + B("1 thrall\nSkeletal Thrall | HP 4 | Morale - | Claw/Bite d4"),
}
