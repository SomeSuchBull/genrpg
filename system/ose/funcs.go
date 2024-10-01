package ose

import (
	"fmt"

	"github.com/genrpg/utils"
)

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
		fmt.Printf("%03d: %s\n", i+1, roomContents)
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

func getTrap() string {
	return "Trap"
}
