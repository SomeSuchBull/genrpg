package shadowdark

import (
	"fmt"
	"math/rand"

	"github.com/genrpg/utils"
)

func getTrap() string {
	var trap string
	tables := [3]string{"Trap", "Trigger", "Effect"}
	for i, v := range traps {
		roll := rand.Intn(len(v))
		trap += fmt.Sprintf("%s: %s ", tables[i], v[roll])
	}
	return trap
}

var engine = map[int]func() string{
	1:  getEmpty,
	2:  getEmpty,
	3:  getTrap,
	4:  getMinorHazard,
	5:  getSoloMonster,
	6:  getNPC,
	7:  getMonsterMob,
	8:  getMajorHazard,
	9:  getTreasure,
	10: getBossMonster,
}

func Stocking(rooms int64) {
	for i := int64(0); i < rooms; i++ {
		roomContents := ""
		roll := utils.TableDie(6)
		f := engine[roll]
		roomContents += f()
		fmt.Printf("%3d: \n", i+1)
	}
}

func getTreasure() string {
	return "Treasure"
}

func getEmpty() string {
	return "Empty"
}

func getMonster() string {
	return "Monster"
}

func getSpecial() string {
	return "Special"
}

func getMinorHazard() string {
	return "Minor Hazard"
}
func getSoloMonster() string {
	return "Solo Monster"
}
func getNPC() string {
	return "NPC"
}
func getMonsterMob() string {
	return "Monster Mob"
}
func getMajorHazard() string {
	return "Major Hazard"
}
func getBossMonster() string {
	return "Boss Monster"
}
