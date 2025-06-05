package shadowdark

import (
	"fmt"

	"github.com/genrpg/utils"
)

type Fighter struct {
	pc                          *PlayerCharacter
	name, armor, weapons, title string
	talents                     []ClassTalent
	features                    []ClassFeature
	hpDie, lvl                  int
}

func (pc *PlayerCharacter) NewFighter() {
	fighter := &Fighter{
		pc:      pc,
		name:    "Fighter",
		armor:   "All armor and shields",
		weapons: "All weapons",
		hpDie:   8,
		lvl:     pc.Lvl,
		features: []ClassFeature{
			ClassFeature(fmt.Sprintf(
				"%s Add your Constitution modifier, if positive, to your gear slots. %s",
				utils.B("Hauler."), utils.I("Added already."))),
			ClassFeature(fmt.Sprintf(
				"%s Choose one type of weapon, such as longswords. You gain +1 to attack and damage with that weapon type. In addition, add half your level to these rolls (round down).",
				utils.B("Weapon Mastery."))),
			ClassFeature(fmt.Sprintf("%s Choose Strength or Dexterity. You have advantage on checks of that type to overcome an opposing force, such as kicking open a stuck door (Strength) or slipping free of rusty chains (Dexterity).",
				utils.B("Grit."))),
		},
	}
	fighter.RollTalent()
	fighter.SetTitle()
	if pc.Stats.Constitution > 10 {
		pc.ItemSlots += pc.Stats.Constitution.Mod()
	}
	pc.Class = fighter
}

func (f *Fighter) Description() string { return "" }

func (f *Fighter) Name() string { return f.name }

func (f *Fighter) Lvl() int { return f.lvl }

func (f *Fighter) Title(string) string { return f.title }

func (f *Fighter) Features() []ClassFeature { return f.features }

func (f *Fighter) Talents() []ClassTalent { return f.talents }

func (f *Fighter) HPDie() int { return f.hpDie }

func (f *Fighter) Spellcaster() bool { return false }

// func (f *Fighter) Assign(pc *PlayerCharacter) { f.pc = pc }

func (f *Fighter) RollTalent() {
	switch utils.D(6, 2) {
	case 2:
		f.talents = append(f.talents, "Gain Weapon Mastery with one additional weapon type")
	case 3, 4, 5, 6:
		f.talents = append(f.talents, ClassTalent(fmt.Sprintf("+1 to melee and ranged attacks %s", utils.I("added already"))))
		f.pc.Attacks.Melee++
		f.pc.Attacks.Ranged++
	case 7, 8, 9:
		f.talents = append(f.talents, ClassTalent(fmt.Sprintf("+2 to Strength, Dexterity, or Constitution stat %s", utils.I("added already"))))
		if f.pc.Optimized {
			// TODO: Implement this
		} else {
			switch utils.D(3) {
			case 1:
				f.pc.Stats.Strength += 2
			case 2:
				f.pc.Stats.Dexterity += 2
			case 3:
				f.pc.Stats.Constitution += 2
			}
		}
	case 10, 11:
		f.talents = append(f.talents, "Choose one kind of armor. You get +1 AC from that armor")
	default:
		f.talents = append(f.talents, "Choose a talent or +2 points to distribute to stats")
	}
}

func (f *Fighter) SetTitle() {
	f.title = fighterTitles[f.pc.Alignment][(f.lvl-1)/2]
}

var fighterTitles = map[Alignment][]string{
	Lawful:  {"Squire", "Cavalier", "Knight", "Thane", "Lord/Lady"},
	Neutral: {"Warrior", "Barbarian", "Battlerager", "Warchief", "Chieftain"},
	Chaotic: {"Knave", "Bandit", "Slayer", "Reaver", "Warlord"},
}
