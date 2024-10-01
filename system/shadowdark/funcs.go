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
	0: getTreasure,
	1: getTreasure,
	2: getMonster,
	3: getMonster,
	4: getSpecial,
	5: getTrap,
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

func getMonster() string {
	return "Monster"
}

func getSpecial() string {
	return "Special"
}
