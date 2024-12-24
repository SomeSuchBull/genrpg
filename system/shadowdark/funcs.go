package shadowdark

import (
	"fmt"
	"math/rand"

	"github.com/genrpg/utils"
)

var engine = []func(int) string{getEmpty, getEmpty, getTrap, getMinorHazard, getSoloMonster, getNPC, getMonsterMob, getMajorHazard, getTreasure, getBossMonster}

func Stocking(rooms int64) {
	for i := 1; i < int(rooms)+1; i++ {
		roomContents := ""
		roll := utils.TableDie(10)
		f := engine[roll]
		roomContents += f(i)
		fmt.Printf("%3d: %s\n", i, roomContents)
	}
}

func getTrap(level int) string {
	var trap string
	tables := [3]string{"Trap", "Trigger", "Effect"}
	for i, v := range traps {
		roll := rand.Intn(len(v))
		trap += fmt.Sprintf("%s: %s", tables[i], v[roll])
		if i < len(traps)-1 {
			trap += " | "
		}
	}
	return trap
}

func getTreasure(level int) string {
	return "Treasure"
}

func getEmpty(level int) string {
	return "Empty"
}

func getMonster(level int) string {
	return "Monster"
}

func getSpecial(level int) string {
	return "Special"
}

func getMinorHazard(level int) string {
	return "Minor Hazard"
}
func getSoloMonster(level int) string {
	return "Solo Monster"
}
func getNPC(level int) string {
	return "NPC"
}
func getMonsterMob(level int) string {
	return "Monster Mob"
}
func getMajorHazard(level int) string {
	return "Major Hazard"
}
func getBossMonster(level int) string {
	return "Boss Monster"
}
