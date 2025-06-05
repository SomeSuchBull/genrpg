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

type Teamster struct{}

func (t *Teamster) Skills() []Skill {
	skills := []Skill{IndustrialEquipment, ZeroG}
	skillNames := []string{IndustrialEquipment.Name, ZeroG.Name}
	availableExpertSkills := append(IndustrialEquipment.PrerequisiteFor, ZeroG.PrerequisiteFor...)
	pickedSkill := Skill{}
	for (pickedSkill.Name == "") || slices.Contains(skillNames, pickedSkill.Name) {
		pickedSkill = TrainedSkills[utils.TD(len(TrainedSkills))]
	}
	skills = append(skills, pickedSkill)
	availableExpertSkills = append(availableExpertSkills, pickedSkill.PrerequisiteFor...)
	skills = append(skills, availableExpertSkills[utils.TD(len(availableExpertSkills))])
	return skills
}

func (t *Teamster) Name() string {
	return "Teamster"
}

func (t *Teamster) Adjustments() (Stats, Saves, int) {
	return Stats{Strength: 5, Speed: 5, Intellect: 5, Combat: 5}, Saves{Sanity: 10, Fear: 10, Body: 10}, 0
}

func (t *Teamster) Loadout() []Equipment {
	return nil
}

func NewTeamster() *Teamster {
	return &Teamster{}
}
