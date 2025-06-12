package mothership

import (
	"fmt"
	"strings"

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
	Loadout() []ItemInterface
}

type Skills []Skill

func (s Skills) String() string {
	value := ""
	trained := []string{}
	expert := []string{}
	master := []string{}
	for _, skill := range s {
		switch skill.Type {
		case SkillTypeTrained:
			trained = append(trained, skill.Name)
		case SkillTypeExpert:
			expert = append(expert, skill.Name)
		case SkillTypeMaster:
			master = append(master, skill.Name)
		}
	}
	if len(trained) > 0 {
		value += fmt.Sprintf("Trained Skills (+10): %s\n", strings.Join(trained, ", "))
	}
	if len(expert) > 0 {
		value += fmt.Sprintf("Expert Skills (+15): %s\n", strings.Join(expert, ", "))
	}
	if len(master) > 0 {
		value += fmt.Sprintf("Master Skills (+20): %s\n", strings.Join(master, ", "))
	}
	return value
}

type Equipment []ItemInterface

func (e Equipment) String() string {
	var value string
	if len(e) > 0 {
		eList := make([]string, len(e))
		for i, item := range e {
			eList[i] = item.String()
		}
		value = fmt.Sprintf("Equipment:\n- %s", strings.Join(eList, "\n- "))
	}
	return value
}

type PC struct {
	Stats       Stats
	Saves       Saves
	Class       Class
	Skills      Skills
	Health      int
	Wounds      int
	Stress      int
	ArmorPoints int
	Equipment   Equipment
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

Health: %d
Wounds: %d
Stress: %d

Armor Points: %d

Skills:
%v
%v

Credits: %d

For details on Equipment, check Player's Survival Guide.`, pc.Class.Name(),
		// Stats
		pc.Stats.Strength,
		pc.Stats.Speed,
		pc.Stats.Intellect,
		pc.Stats.Combat,
		// Saves
		pc.Saves.Sanity,
		pc.Saves.Fear,
		pc.Saves.Body,

		pc.Health, pc.Wounds, pc.Stress, pc.ArmorPoints, pc.Skills, pc.Equipment, pc.Credits)
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

func (pc *PC) Loadout() {
	pc.Equipment = pc.Class.Loadout()
	for _, item := range pc.Equipment {
		if armor, ok := item.(Armor); ok {
			pc.ArmorPoints = armor.ArmorPoints
		}
	}
}
