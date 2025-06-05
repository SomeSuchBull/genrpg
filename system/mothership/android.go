package mothership

import (
	"slices"

	"github.com/genrpg/utils"
)

// type Class interface {
// 	Skills()
// 	Name() string
// 	Adjustments()
// 	Loadout()
// }

type Android struct{}

func (a *Android) Skills() []Skill {
	skills := []Skill{Linguistics, Computers, Mathematics}
	skillNames := []string{Linguistics.Name, Computers.Name, Mathematics.Name}
	availableExpertSkills := append(append(Linguistics.PrerequisiteFor, Computers.PrerequisiteFor...), Mathematics.PrerequisiteFor...)
	if utils.D(2) == 1 {
		skills = append(skills, availableExpertSkills[utils.TD(len(availableExpertSkills))])
	} else {
		for range 2 {
			pickedSkill := Skill{}
			for (pickedSkill.Name == "") || slices.Contains(skillNames, pickedSkill.Name) {
				pickedSkill = TrainedSkills[utils.TD(len(TrainedSkills))]
			}
			skills = append(skills, pickedSkill)
			skillNames = append(skillNames, pickedSkill.Name)
		}
	}
	return skills
}

func (a *Android) Name() string {
	return "Android"
}

func (a *Android) Adjustments() (Stats, Saves, int) {
	stats := Stats{Intellect: 20}
	switch utils.D(3) {
	case 1:
		stats.Strength = -10
	case 2:
		stats.Speed = -10
	case 3:
		stats.Combat = -10
	}
	return stats, Saves{Fear: 60}, 1
}

func (a *Android) Loadout() []Equipment {
	return nil
}

func NewAndroid() *Android {
	return &Android{}
}
