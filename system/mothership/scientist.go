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

type Scientist struct{}

func (s *Scientist) Skills() []Skill {
	masterSkill := MasterSkills[utils.TD(len(MasterSkills))]
	expertSkill := masterSkill.Prerequisites[utils.TD(len(masterSkill.Prerequisites))]
	trainedSkill := TrainedSkills[utils.TD(len(TrainedSkills))]
	skills := []Skill{masterSkill, expertSkill, trainedSkill}
	skillNames := []string{masterSkill.Name, expertSkill.Name, trainedSkill.Name}
	pickedSkill := Skill{}
	for (pickedSkill.Name == "") || slices.Contains(skillNames, pickedSkill.Name) {
		pickedSkill = TrainedSkills[utils.TD(len(TrainedSkills))]
	}
	skills = append(skills, pickedSkill)

	return skills
}

func (s *Scientist) Name() string {
	return "Scientist"
}

func (s *Scientist) Adjustments() (Stats, Saves, int) {
	stats := Stats{Intellect: 10}
	switch utils.D(3) {
	case 1:
		stats.Strength = 5
	case 2:
		stats.Speed = 5
	case 3:
		stats.Combat = 5
	}
	return stats, Saves{Sanity: 30}, 0
}

func (s *Scientist) Loadout() []ItemInterface {
	return ScientistLoadouts[utils.TD(len(ScientistLoadouts))]
}

func NewScientist() *Scientist {
	return &Scientist{}
}
