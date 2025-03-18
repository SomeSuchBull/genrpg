package shadowdark

import (
	"fmt"

	"github.com/genrpg/utils"
)

type Thief struct {
	pc                          *PlayerCharacter
	name, armor, weapons, title string
	talents                     []ClassTalent
	features                    []ClassFeature
	hpDie, lvl                  int
}

func (pc *PlayerCharacter) NewThief() {
	thief := &Thief{
		pc:      pc,
		name:    "Thief",
		armor:   "Leather armor, mithral chainmail",
		weapons: "Club, crossbow, dagger, shortbow, shortsword",
		hpDie:   4,
		lvl:     pc.Lvl,
		features: []ClassFeature{
			ClassFeature(fmt.Sprintf(
				"%s If you hit a creature who is unaware of your attack, you deal an extra weapon die of damage. Add additional weapon dice of damage equal to half your level (round down).",
				utils.B("Backstab."),
			)),
			ClassFeature(fmt.Sprintf(
				"%s You are adept at thieving skills and have the necessary tools of the trade secreted on your person (they take up no gear slots).",
				utils.B("Thievery."),
			)),
			ClassFeature("You are trained in the following tasks and have advantage on any associated checks:" +
				"• Climbing\n• Sneaking and hiding\n• Applying disguises\n• Finding and disabling traps\n• Delicate tasks such as picking pockets and opening locks"),
		},
	}
	thief.RollTalent()
	thief.SetTitle()
	pc.Class = thief
}

func (t *Thief) Description() string { return "" }

func (t *Thief) Name() string { return t.name }

func (t *Thief) Lvl() int { return t.lvl }

func (t *Thief) Title(string) string { return t.title }

func (t *Thief) Features() []ClassFeature { return t.features }

func (t *Thief) Talents() []ClassTalent { return t.talents }

func (t *Thief) HPDie() int { return t.hpDie }

func (t *Thief) Spellcaster() bool { return false }

// func (t *Thief) Assign(pc *PlayerCharacter) { t.pc = pc }

func (t *Thief) RollTalent() {
	switch utils.D(6) + utils.D(6) {
	case 2:
		t.talents = append(t.talents, "Gain advantage on initiative rolls (reroll if duplicate)")
	case 3, 4, 5:
		t.talents = append(t.talents, "Your Backstab deals +1 dice of damage")
	case 6, 7, 8, 9:
		t.talents = append(t.talents, ClassTalent(fmt.Sprintf("+2 to Strength, Dexterity, or Charisma stat %s", utils.I("added already"))))
		if t.pc.Optimized {
			// TODO: Implement this
		} else {
			switch utils.D(3) {
			case 1:
				t.pc.Stats.Strength += 2
			case 2:
				t.pc.Stats.Dexterity += 2
			case 3:
				t.pc.Stats.Charisma += 2
			}
		}
	case 10, 11:
		t.talents = append(t.talents, ClassTalent(fmt.Sprintf("+1 to melee and ranged attacks %s", utils.I("added already"))))
		t.pc.Attacks.Melee++
		t.pc.Attacks.Ranged++
	default:
		t.talents = append(t.talents, "Choose a talent or +2 points to distribute to stats")

	}
}

func (t *Thief) SetTitle() {
	t.title = thiefTitles[t.pc.Alignment][(t.lvl-1)/2]
}

var thiefTitles = map[Alignment][]string{
	Lawful:  {"Footpad", "Burglar", "Rook", "Underboss", "Boss"},
	Neutral: {"Robber", "Outlaw", "Rogue", "Renegade", "Bandit King/Queen"},
	Chaotic: {"Thug", "Cutthroat", "Shadow", "Assassin", "Wraith"},
}
