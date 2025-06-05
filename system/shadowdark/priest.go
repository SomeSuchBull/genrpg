package shadowdark

import (
	"fmt"

	"github.com/genrpg/utils"
)

type Priest struct {
	pc                          *PlayerCharacter
	name, armor, weapons, title string
	talents                     []ClassTalent
	features                    []ClassFeature
	hpDie, lvl                  int
}

func (pc *PlayerCharacter) NewPriest() {
	priest := &Priest{
		pc:      pc,
		name:    "Priest",
		armor:   "All armor and shields",
		weapons: "Club, crossbow, dagger, mace, longsword, staff, warhammer",
		hpDie:   6,
		lvl:     pc.Lvl,
		features: []ClassFeature{
			ClassFeature(fmt.Sprintf(
				"%s You know Celestial, Diabolic, or Primordial.",
				utils.B("Languages."),
			)),
			ClassFeature(fmt.Sprintf(
				"%s You know the %s spell. It doesn't count toward your number of known spells.",
				utils.B("Turn Undead."), utils.I("turn undead"),
			)),
			ClassFeature(fmt.Sprintf(
				"%s Choose a god to serve who matches your alignment (see Deities, pg. 28). You have a holy symbol for your god (it takes up no gear slots).",
				utils.B("Deity."),
			)),
			ClassFeature(fmt.Sprintf(
				"%s You can cast priest spells you know.\nYou know two tier 1 spells of your choice from the priest spell list on pg. 51.\nEach time you gain a level, you choose new priest spells to learn according to the Priest Spells Known table.\nFor casting priest spells, see Spellcasting on pg. 44.",
				utils.B("Spellcasting."),
			)),
		},
	}
	priest.RollTalent()
	priest.SetTitle()
	pc.Class = priest
}

func (p *Priest) Description() string { return "" }

func (p *Priest) Name() string { return p.name }

func (p *Priest) Lvl() int { return p.lvl }

func (p *Priest) Title(string) string { return p.title }

func (p *Priest) Features() []ClassFeature { return p.features }

func (p *Priest) Talents() []ClassTalent { return p.talents }

func (p *Priest) HPDie() int { return p.hpDie }

func (p *Priest) Spellcaster() bool { return false }

// func (p *Priest) Assign(pc *PlayerCharacter) { p.pc = pc }

func (p *Priest) RollTalent() {
	switch utils.D(6, 2) {
	case 2:
		p.talents = append(p.talents, "Gain advantage on casting one spell you know")
	case 3, 4, 5, 6:
		p.talents = append(p.talents, ClassTalent(fmt.Sprintf("+1 to melee or ranged attacks %s", utils.I("added already"))))
		if p.pc.Optimized {
		} else {
			switch utils.D(2) {
			case 1:
				p.pc.Attacks.Melee++
			case 2:
				p.pc.Attacks.Ranged++
			}
		}
	case 7, 8, 9:
		p.talents = append(p.talents, ClassTalent(fmt.Sprintf("+1 to priest spellcasting checks %s", utils.I("added already"))))
		p.pc.Attacks.Cast++
	case 10, 11:
		p.talents = append(p.talents, ClassTalent(fmt.Sprintf("+2 to Strength or Wisdom stat %s", utils.I("added already"))))
		if p.pc.Optimized {
		} else {
			switch utils.D(2) {
			case 1:
				p.pc.Stats.Strength += 2
			case 2:
				p.pc.Stats.Wisdom += 2
			}
		}
	default:
		p.talents = append(p.talents, "Choose a talent or +2 points to distribute to stats")

	}
}

func (p *Priest) SetTitle() {
	p.title = priestTitles[p.pc.Alignment][(p.lvl-1)/2]
}

var priestTitles = map[Alignment][]string{
	Lawful:  {"Acolyte", "Crusader", "Templar", "Champion", "Paladin"},
	Neutral: {"Seeker", "Invoker", "Haruspex", "Mystic", "Oracle"},
	Chaotic: {"Initiate", "Zealot", "Cultist", "Scourge", "Chaos Knight"},
}
