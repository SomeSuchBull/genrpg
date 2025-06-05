package shadowdark

import (
	"fmt"

	"github.com/genrpg/utils"
)

type Wizard struct {
	pc                          *PlayerCharacter
	name, armor, weapons, title string
	talents                     []ClassTalent
	features                    []ClassFeature
	hpDie, lvl                  int
}

func (pc *PlayerCharacter) NewWizard() {
	wizard := &Wizard{
		pc:      pc,
		name:    "Wizard",
		armor:   "None",
		weapons: "Dagger, staff",
		hpDie:   6,
		lvl:     pc.Lvl,
		features: []ClassFeature{
			ClassFeature(fmt.Sprintf(
				"%s You know two additional common languages and two rare languages (see pg. 32).",
				utils.B("Languages."),
			)),
			ClassFeature(fmt.Sprintf(
				"%s You can permanently learn a wizard spell from a spell scroll by studying it for a day and succeeding on a DC 15 Intelligence check.\nWhether you succeed or fail, you expend the spell scroll.\nSpells you learn in this way don't count toward your known spells",
				utils.B("Learning Spells."),
			)),
			ClassFeature(fmt.Sprintf(
				"%s You can cast wizard spells you know.\nYou know three tier 1 spells of your choice from the wizard spell list (see pg. 52).\nEach time you gain a level, you choose new wizard spells to learn according to the Wizard Spells Known table.\nFor casting wizard spells, see Spellcasting on pg. 44.",
				utils.B("Spellcasting."),
			)),
		},
	}
	wizard.RollTalent()
	wizard.SetTitle()
	pc.Class = wizard
}

func (w *Wizard) Description() string { return "" }

func (w *Wizard) Name() string { return w.name }

func (w *Wizard) Lvl() int { return w.lvl }

func (w *Wizard) Title(string) string { return w.title }

func (w *Wizard) Features() []ClassFeature { return w.features }

func (w *Wizard) Talents() []ClassTalent { return w.talents }

func (w *Wizard) HPDie() int { return w.hpDie }

func (w *Wizard) Spellcaster() bool { return false }

// func (w *Wizard) Assign(pc *PlayerCharacter) { w.pc = pc }

func (w *Wizard) RollTalent() {
	switch utils.D(6, 2) {
	case 2:
		w.talents = append(w.talents, "Make 1 random magic item of a type you choose (pg. 282)")
	case 3, 4, 5, 6, 7:
		w.talents = append(w.talents, ClassTalent(fmt.Sprintf("+2 to Intelligence stat or +1 to wizard spellcasting checks %s", utils.I("added already"))))
		if w.pc.Stats.Intelligence < 18 {
			w.pc.Stats.Intelligence += 2
		} else {
			w.pc.Attacks.Cast++
		}
	case 8, 9:
		w.talents = append(w.talents, ClassTalent("Gain advantage on casting one spell you know"))
	case 10, 11:
		w.talents = append(w.talents, ClassTalent("Learn one additional wizard spell of any tier you know"))
	default:
		w.talents = append(w.talents, "Choose a talent or +2 points to distribute to stats")

	}
}

func (w *Wizard) SetTitle() {
	w.title = wizardTitles[w.pc.Alignment][(w.lvl-1)/2]
}

var wizardTitles = map[Alignment][]string{
	Lawful:  {"Acolyte", "Crusader", "Templar", "Champion", "Paladin"},
	Neutral: {"Seeker", "Invoker", "Haruspex", "Mystic", "Oracle"},
	Chaotic: {"Initiate", "Zealot", "Cultist", "Scourge", "Chaos Knight"},
}
