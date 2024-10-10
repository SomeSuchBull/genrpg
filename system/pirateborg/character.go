package pirateborg

import "github.com/genrpg/utils"

type Stats struct {
	Strength  int `json:"strength"`
	Agility   int `json:"agility"`
	Presence  int `json:"presence"`
	Toughness int `json:"toughness"`
	Spirit    int `json:"spirit"`
}

type PlayerCharacter struct {
	Stats    Stats       `json:"stats"`
	Class    PlayerClass `json:"class"`
	HP       int         `json:"hp"`
	Name     string      `json:"name"`
	Nickname string      `json:"nickname"`
	Weapons  []Weapon    `json:"weapons"`
	Clothing Clothing    `json:"clothing"`
	Armor    string      `json:"armor"`
	Hat      Hat         `json:"hat"`
	Gear     []Item      `json:"gear"`
}

type PlayerClass interface {
	Level() int
	LevelUp()
	String() string
	Description() string
	GetItems() []Item
	GetWeapons() []Weapon
}

type Feature interface {
	String() string
}

func NewCharacter(additionalClasses ...bool) *PlayerCharacter {
	pc := &PlayerCharacter{}
	pc.GetClass()
	pc.GetStats()
	return pc
}

func (pc *PlayerCharacter) GetClass(additionalClasses ...bool) {
	res := utils.D(6)
	if len(additionalClasses) > 0 && additionalClasses[0] {
		res = utils.D(8)
	}
	switch res {
	case 1:
		pc.NewBrute()
	}
}

func (pc *PlayerCharacter) GetStats() {
	pc.GetStartingStats()
	// TODO: Class mods on stats
}

func (pc *PlayerCharacter) GetStartingStats() {
	// TODO: determine stats
}
