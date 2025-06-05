package mothership

import (
	"fmt"

	"github.com/genrpg/utils"
)

type Stats struct {
	Strength  int
	Speed     int
	Intellect int
	Combat    int
}

type Saves struct {
	Sanity int
	Fear   int
	Body   int
}

type Class interface {
	Skills() []Skill
	Name() string
	Adjustments() (Stats, Saves, int)
	Loadout() []Equipment
}

type Skills map[string]int

type PC struct {
	Stats       Stats
	Saves       Saves
	Class       Class
	Skills      []Skill
	Health      int
	Wounds      int
	Stress      int
	ArmorPoints int
	Equipment   []Equipment
	Credits     int
}

func NewPC() *PC {
	pc := &PC{
		Stats: Stats{
			Strength:  25 + utils.D(10, 2),
			Speed:     25 + utils.D(10, 2),
			Intellect: 25 + utils.D(10, 2),
			Combat:    25 + utils.D(10, 2),
		},
		Saves: Saves{
			Sanity: 10 + utils.D(10, 2),
			Fear:   10 + utils.D(10, 2),
			Body:   10 + utils.D(10, 2),
		},
		Health:  10 + utils.D(10),
		Wounds:  2,
		Stress:  2,
		Credits: (utils.D(10, 2)) * 10,
	}
	pc.Class = GetClass()
	pc.GetSkills()
	pc.AdjustNumbers()
	pc.Loadout()
	fmt.Println(pc)
	return pc
}

func (pc *PC) String() string {
	value := fmt.Sprintf(`
Class: %s
Stats:
  Strength:  %d
  Speed:     %d
  Intellect: %d
  Combat:    %d
Saves:
  Sanity: %d
  Fear:   %d
  Body:   %d
Health:      %d
Wounds:      %d
Stress:      %d
Armor Points: %d
Skills: %v
Equipment: %v
Credits:     %d
`, pc.Class.Name(),
		pc.Stats.Strength,
		pc.Stats.Speed,
		pc.Stats.Intellect,
		pc.Stats.Combat,
		pc.Saves.Sanity,
		pc.Saves.Fear,
		pc.Saves.Body,
		pc.Health,
		pc.Wounds,
		pc.Stress,
		pc.ArmorPoints,
		pc.Skills,
		pc.Equipment,
		pc.Credits)
	return value
}

func GetClass() Class {
	switch utils.D(4) {
	default:
		return NewMarine()
	case 2:
		return NewAndroid()
	case 3:
		return NewScientist()
	case 4:
		return NewTeamster()
	}
}

func (pc *PC) GetSkills() {
	pc.Skills = pc.Class.Skills()
}

func (pc *PC) AdjustNumbers() {
	stats, saves, wounds := pc.Class.Adjustments()

	pc.Stats.Strength += stats.Strength
	pc.Stats.Speed += stats.Speed
	pc.Stats.Intellect += stats.Intellect
	pc.Stats.Combat += stats.Combat

	pc.Saves.Sanity += saves.Sanity
	pc.Saves.Fear += saves.Fear
	pc.Saves.Body += saves.Body

	pc.Wounds += wounds
}

func (pc *PC) Loadout() {}
