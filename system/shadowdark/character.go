package shadowdark

import (
	"fmt"

	"github.com/genrpg/utils"
)

type PCStat int

func (s PCStat) Mod() int {
	switch s {
	case 4, 5:
		return -3
	case 6, 7:
		return -2
	case 8, 9:
		return -1
	case 10, 11:
		return 0
	case 12, 13:
		return 1
	case 14, 15:
		return 2
	case 16, 17:
		return 3
	default:
		switch {
		case s <= 3:
			return -4
		default:
			return 4
		}
	}
}

type PCStats struct {
	Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma PCStat
}

type Attacks struct {
	Melee, MeleeDamage, Ranged, RangedDamage, Cast int
}

type ClassFeature string
type ClassTalent string
type Background string

type PlayerClass interface {
	Description() string
	Name() string
	Lvl() int
	Title(string) string
	Features() []ClassFeature
	Talents() []ClassTalent
	RollTalent()
	HPDie() int
	Spellcaster() bool
	// Assign(*PlayerCharacter)
}

type PlayerCharacter struct {
	Stats      PCStats
	Attacks    Attacks
	Ancestry   Ancestry
	Class      PlayerClass
	HP         int
	Lvl        int
	Background Background
	Alignment  Alignment
	Name       string
	// TODO: change to Gear struct, which has to be created
	ItemSlots        int
	Gear             []string
	Optimized, Extra bool
}

func (pc *PlayerCharacter) GetAncestry() {}

func (pc *PlayerCharacter) GetClass() {
	if !pc.Optimized {
		switch utils.D(4) {
		default:
			pc.NewFighter()
		case 2:
			pc.NewThief()
			// case 3:
			// 	pc.NewCleric()
			// case 4:
			// 	pc.NewMage()
		}
	}
}

func (pc *PlayerCharacter) RollHP() {
	if pc.HP == 0 {
		pc.HP = max(utils.D(pc.Class.HPDie())+pc.Stats.Constitution.Mod(), 1)
	}
}

func (pc *PlayerCharacter) SetItemSlots() {
	pc.ItemSlots += max(int(pc.Stats.Strength), 10)
}

func (pc *PlayerCharacter) GetAlignment() {
	switch utils.D(3) {
	case 1:
		pc.Alignment = Lawful
	case 2:
		pc.Alignment = Neutral
	case 3:
		pc.Alignment = Chaotic
	}
}

func (pc *PlayerCharacter) GetBackground() {
	pc.Background = Background(backgrounds[utils.TD(len(backgrounds))])
}

func (pc *PlayerCharacter) GetGear() {}

func GenerateCharacter(optimized, extra bool) {
	pc := NewCharacter(optimized, extra)
	fmt.Println(pc)
}

func NewCharacter(optimized, extra bool) PlayerCharacter {
	pc := &PlayerCharacter{Optimized: optimized, Extra: extra}
	pc.Stats = rollPCStats()
	pc.GetAlignment()
	pc.GetClass()
	pc.GetAncestry()
	pc.Ancestry.FeatureFunc(pc)
	pc.RollHP()
	pc.SetItemSlots()
	pc.GetBackground()
	pc.GetGear()
	return *pc
}

func rollPCStats() PCStats {
	stats := make([]PCStat, 6)
	statsGood := false
	for !statsGood {
		for i := range stats {
			stat := utils.D(6) * utils.D(6) * utils.D(6)
			stats[i] = PCStat(stat)
			if stat >= 14 {
				statsGood = true
			}
		}
	}
	return PCStats{stats[0], stats[1], stats[2], stats[3], stats[4], stats[5]}
}
