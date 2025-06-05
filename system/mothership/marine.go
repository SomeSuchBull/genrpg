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

type Marine struct{}

func (m *Marine) Skills() []Skill {
	skills := []Skill{MilitaryTraining, Athletics}
	skillNames := []string{MilitaryTraining.Name, Athletics.Name}
	availableExpertSkills := MilitaryTraining.PrerequisiteFor
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

func (m *Marine) Name() string {
	return "Marine"
}

func (m *Marine) Adjustments() (Stats, Saves, int) {
	return Stats{Combat: 10}, Saves{Fear: 20, Body: 10}, 1
}

func (m *Marine) Loadout() []Equipment {
	return nil
}

func NewMarine() *Marine {
	return &Marine{}
}
